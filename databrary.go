/*
Package databrary is the REST API backend for http://nyu.databrary.org/


This is a simple REST API backend that connects the frontend to the data layer
(currently Postgres and Redis). It uses

https://github.com/go-chi/chi for routing

https://github.com/databrary/sqlboiler for models (a fork of https://github.com/vattle/sqlboiler)

https://github.com/databrary/scs for sessions (a fork of https://github.com/alexedwards/scs)

and various other utils (logging/mail) and middleware.
*/
package main

import (
	"net/http"

	"os"
	"path/filepath"

	"fmt"
	"github.com/databrary/databrary-backend-go/config"
	"github.com/databrary/databrary-backend-go/db"
	"github.com/databrary/databrary-backend-go/logging"
	"github.com/databrary/databrary-backend-go/routes"
	"github.com/databrary/databrary-backend-go/services/redis"
	"github.com/databrary/databrary-backend-go/services/sessions"
	"github.com/databrary/sqlboiler/boil"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"gopkg.in/alecthomas/kingpin.v2"
	"regexp"
	"strings"
	"time"
)

var (
	projRoot   string
	configPath *string
)

// initialize all of the auxiliary services
func init() {
	goPaths := strings.Split(filepath.Join(os.Getenv("GOPATH"), "src/github.com/databrary/databrary-backend-go/"), ":")
	// if there's a vendor directory then there will be two gopaths (that's how vendoring works). we want the second one
	// which is the actual gopath
	if len(goPaths) == 2 {
		projRoot = goPaths[1]
	} else if len(goPaths) == 1 {
		projRoot = goPaths[0]
	} else {
		panic(fmt.Sprintf("unexpected gopath %#v", goPaths))
	}

	configPath = kingpin.Flag("config", "Path to config file").
		Default(filepath.Join(projRoot, "config/databrary_dev.toml")).
		Short('c').
		String()

	// parse command line flags
	kingpin.Version("0.0.0")
	kingpin.Parse()

	if configPath, err := filepath.Abs(*configPath); err != nil {
		panic("command line config file path error")
	} else {
		log.InitLgr(config.InitConf(configPath))
	}

	// initialize db connection
	err := db.InitDB(config.GetConf())
	if err != nil {
		panic(err.Error())
	}

	// initialize redis connection
	redis.InitRedisStore(config.GetConf())

	if config.GetConf().GetString("log.level") == "DEBUG" {
		// print to stdout the sql generated by sqlboiler
		boil.DebugMode = true
	}
}

func main() {
	conf := config.GetConf()
	// most of these i just got from somewhere that claimed they were sane defaults
	// they should be explored in depth
	secureMiddleware := secure.New(secure.Options{
		// i don't remember if the ports need to be included. i think they don't
		AllowedHosts:          []string{"localhost:3000", "localhost:3444"},
		HostsProxyHeaders:     []string{"X-Forwarded-Host"},
		SSLRedirect:           true,
		SSLHost:               "",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
		IsDevelopment:         true,
	})

	r := chi.NewRouter()
	// again standard stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(log.NewStructuredLogger(log.Logger))
	r.Use(middleware.Recoverer)
	r.Use(secureMiddleware.Handler)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.StripSlashes)
	r.Use(sessions.NewSessionManager())

	// rate limiter middlware
	rateLimiter, err := routes.NewRateLimiter()
	if err != nil {
		log.WrapErrLogFatal(err, "couldn't create rate limiter")
	}
	r.Use(rateLimiter.RateLimit)

	// cross origin resource sharing. in production this should be the hostname
	// of wherever the frontend is served from
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: []string{"http://localhost:3000", "https://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		Debug:            true,
	})
	r.Use(c.Handler)

	// all api routes are mounted under /api
	r.Mount("/api", routes.Api())

	// root should host index once frontend is done
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nothing here yet"))
	})

	// assets (js and svg and such
	r.FileServer("/public", http.Dir("public"))

	// generate the api docs (stored in api.md)
	GenerateAPI(r)

	// the only reason that there's an address config is because static emails are generated with that address
	// embedded in them
	addr := conf.GetString("address.domain") + ":" + conf.GetString("address.backend_port")
	fmt.Printf("serving on %s://%s/\n", conf.GetString("address.scheme"), addr)

	certPath := conf.GetString("ssl.cert")
	keyPath := conf.GetString("ssl.key")
	err = http.ListenAndServeTLS(addr, certPath, keyPath, r)
	panic(err)
}

func GenerateAPI(r chi.Router) {
	m := docgen.MarkdownOpts{
		ProjectPath:        "github.com/databrary/databrary-backend-go",
		Intro:              "Databrary 2.0 API",
		ForceRelativeLinks: true,
	}
	// skip documenting middleware
	re := regexp.MustCompile(`^- \[`)
	f, _ := os.Create(filepath.Join(projRoot, "api.md"))
	defer f.Close()
	docs := docgen.MarkdownRoutesDoc(r, m)
	for _, v := range strings.Split(docs, "\n") {
		if re.FindStringIndex(v) == nil {
			f.WriteString(v + "\n")
		}
	}
	f.Sync()
}
