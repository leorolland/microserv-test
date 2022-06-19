# Microserv
A small cloud-native app that connects to a PostgreSQL instance and exposes a web service.

## Getting started
Using Docker
```sh
# With default settings
docker run leorolland/microserv-test:v0.3.0
# With .env file
docker run --env-file .env leorolland/microserv-test:v0.3.0
```
Build & run using vars in `.env` (copy `.env.sample` to `.env` before) (requires Go compiler)
```sh
make build && env $(cat .env | xargs) ./dist/microserv
```

## PostgreSQL connection

If env var `PG_ENABLED` is set, the app will try to connect to a PostgreSQL instance :

You must also define the following vars :
- `PG_HOST`
- `PG_PORT`
- `PG_USER`
- `PG_PASSWORD`
- `PG_DB`
- `PG_TABLE` (optional)

### Testing PostgreSQL connection (with Docker)
```sh
# Create a docker network in order to make the 2 containers communicate
docker network create postgres
# Start a postgresql container
docker run --name postgres -e POSTGRES_PASSWORD=postgres --network postgres -p 5432:5432 -d postgres
# Start a microserv-test container and connect to postgresql instance
docker run \
-e PG_ENABLED=true \
-e PG_HOST=postgres \
-e PG_PORT=5432 \
-e PG_USER=postgres \
-e PG_PASSWORD=postgres \
-e PG_DB=postgres \
--network postgres -p 8000:8000 leorolland/microserv-test:v0.3.0
```

## Expose a WebService

If env var `WS_ENABLED` is set, the app will expose a very basic webservice :

You must also define the following vars :
- `WS_PORT`

Exposed endpoints
- GET /
- GET /ready
- GET /healthz
