FROM postgres
ENV POSTGRES_PASSWORD mysecretpassword
RUN apt-get update
RUN apt-get install -y libpq-dev postgresql-server-dev-all git gcc
WORKDIR /usr/local/src
RUN git clone https://github.com/databrary/databrary-backend-go
WORKDIR databrary-backend-go
# TODO: set this to databrary_go eventually
ARG commit=origin/master
RUN git remote update && git checkout $commit
RUN make -C pgranges install

ADD init-user-db.sh /usr/local/src/databrary/init-user-db.sh
RUN chmod 755 /usr/local/src/databrary/init-user-db.sh
ADD remove-superuser-db.sh /usr/local/src/databrary/remove-superuser-db.sh
RUN chmod 755 /usr/local/src/databrary/remove-superuser-db.sh
