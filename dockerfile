FROM postgres:15

COPY initdb/init.sql /docker-entrypoint-initdb.d

ENV POSTGRES_USER = root

ENV POSTGRES_PASSWORD = root

ENV POSTGRES_DB = todo

CMD ["docker-entrypoint.sh", "postgres"]