// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"flag"
	"fmt"
	"github.com/databrary/sqlboiler/boil"
	"github.com/kat-co/vala"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var flagDebugMode = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")

var (
	dbMain *pgTester
)

func TestMain(m *testing.M) {
	if dbMain == nil {
		fmt.Println("no dbMain tester interface was ready")
		os.Exit(-1)
	}

	rand.Seed(time.Now().UnixNano())
	dbNameRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var err error

	// Load configuration
	err = initViper()
	if err != nil {
		fmt.Println("unable to load config file")
		os.Exit(-2)
	}

	setConfigDefaults()
	if err := validateConfig("postgres"); err != nil {
		fmt.Println("failed to validate config", err)
		os.Exit(-3)
	}

	// Set DebugMode so we can see generated sql statements
	flag.Parse()
	boil.DebugMode = *flagDebugMode

	if err = dbMain.setupMain(); err != nil {
		fmt.Println("Unable to execute setupMain:", err)
		os.Exit(-4)
	}

	dbMain.testDbConn, err = dbMain.conn(dbMain.TestDBName)
	if err != nil {
		fmt.Println("failed to get test connection:", err)
	}

	var code int
	boil.SetDB(dbMain.testDbConn)
	code = m.Run()

	if err = dbMain.teardown(); err != nil {
		fmt.Println("Unable to execute test teardown:", err)
		os.Exit(-5)
	}

	os.Exit(code)
}

func initViper() error {
	GOPATH := os.Getenv("GOPATH")
	// don't get smart and change sqlboiler.toml to config.toml or something because you won't be able
	// to regen models using sqlboiler executable
	configPath := filepath.Join(GOPATH, "src/github.com/databrary/databrary-backend-go/config/sqlboiler/sqlboiler.toml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	return err
}

// setConfigDefaults is only necessary because of bugs in viper, noted in main
func setConfigDefaults() {
	if viper.GetString("postgres.sslmode") == "" {
		viper.Set("postgres.sslmode", "require")
	}
	if viper.GetInt("postgres.port") == 0 {
		viper.Set("postgres.port", 5432)
	}
	if viper.GetString("mysql.sslmode") == "" {
		viper.Set("mysql.sslmode", "true")
	}
	if viper.GetInt("mysql.port") == 0 {
		viper.Set("mysql.port", 3306)
	}
}

func validateConfig(driverName string) error {
	if driverName == "postgres" {
		return vala.BeginValidation().Validate(
			vala.StringNotEmpty(viper.GetString("postgres.user"), "postgres.user"),
			vala.StringNotEmpty(viper.GetString("postgres.host"), "postgres.host"),
			vala.Not(vala.Equals(viper.GetInt("postgres.port"), 0, "postgres.port")),
			vala.StringNotEmpty(viper.GetString("postgres.dbname"), "postgres.dbname"),
			vala.StringNotEmpty(viper.GetString("postgres.sslmode"), "postgres.sslmode"),
		).Check()
	}

	if driverName == "mysql" {
		return vala.BeginValidation().Validate(
			vala.StringNotEmpty(viper.GetString("mysql.user"), "mysql.user"),
			vala.StringNotEmpty(viper.GetString("mysql.host"), "mysql.host"),
			vala.Not(vala.Equals(viper.GetInt("mysql.port"), 0, "mysql.port")),
			vala.StringNotEmpty(viper.GetString("mysql.dbname"), "mysql.dbname"),
			vala.StringNotEmpty(viper.GetString("mysql.sslmode"), "mysql.sslmode"),
		).Check()
	}

	return errors.New("not a valid driver name")
}
