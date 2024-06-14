FROM alpine
WORKDIR /migration

ARG MYSQLDEF_URL=https://github.com/sqldef/sqldef/releases/download/v0.17.11/mysqldef_linux_amd64.tar.gz

RUN apk add curl
COPY . .
RUN curl -L -o mysqldef.tar.gz ${MYSQLDEF_URL} \
    && tar -zxf mysqldef.tar.gz \
    && rm mysqldef.tar.gz \
    && mv mysqldef /usr/local/bin/

CMD [ "./migration.sh", "schema.sql" ]
