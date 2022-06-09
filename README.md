# Microserv
A small cloud-native app that connects to a PostgreSQL instance and exposes a web service.

## PostgreSQL

If env var `PG_ENABLED` is set, the app will try to connect to a PostgreSQL instance :

You must also define the following vars :
- `PG_HOST`
- `PG_PORT`
- `PG_USER`
- `PG_PASSWORD`
- `PG_DB`

## WebService

If env var `WS_ENABLED` is set, the app will expose a very basic webservice :

You must also define the following vars :
- `WS_PORT`

Exposed endpoints
- GET /healthz
- GET /ready
- GET /
