FROM postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB world
COPY world.sql /docker-entrypoint-initdb.d/

docker build -t my-postgres-db ./

docker images -a