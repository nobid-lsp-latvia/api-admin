# Api-Admin

Admin API

## Built with Azugo Go Web Framework

This project is built using the [Azugo Go Web Framework](https://azugo.io), a powerful and flexible framework for building modern web applications in Go. Check out the [Azugo GitHub page](https://github.com/azugo) for more information and documentation.

<!-- TOC -->

- [Api-Admin](#repo_name_title)
  - [Development](#development)
    - [Prepare dependecies](#prepare-dependecies)
    - [Local development](#local-development)
    - [Before commit](#before-commit)
  - [Environment variables](#environment-variables)
    - [Local example](#local-example)

<!-- /TOC -->

## Development

### Prepare dependecies

```sh
go mod download
go generate ./...
```

### Local development

To build in VS Code use `Ctrl`+`Shift`+`B`.

To debug project in VS Code use `F5`.

### Before commit

> CI requires linted, formatted code

You should run:

```sh
gofmt -s -w ./..
```

or

```sh
gofumpt -w ./..
```

and fix any errors reported by

```sh
golangci-lint run
```

## Environment variables

In order to run the service you need configure environment variables. List of environment variables:

| Variable | Description | Default value | Required |
| --- | --- | --- | --- |
| `SERVER_URLS` | An server URL or multiple URLS separated by semicolon to listen on. | 0.0.0.0:8080 | Yes |
| `ENVIRONMENT` | Environment name. Possible values: `Development`, `Staging`, `Production` | `Development` | Yes |
| `BASE_PATH` | Base path for all routes | `/` (or take value from `SERVER_URLS` path if exists) | No |
| `ACCESS_LOG_ENABLED` | Enable access log | `true` | Yes |
| `REVERSE_PROXY_LIMIT` | Limit for reverse proxy. | `1` | No |
| `REVERSE_PROXY_TRUSTED_IPS` | List of trusted IP addresses for reverse proxy. Separated by `;` | `"127.0.0.1"` | No |
| `REVERSE_PROXY_TRUSTED_HEADERS` | List of trusted headers for reverse proxy. Separated by `;` | `X-Real-IP; X-Forwarded-For` | No |
| `LOG_LEVEL` | Minimal log level. Allowed values are `debug`, `info`, `warn`, `error`, `fatal`, `panic` | `info` | Yes |
| `CACHE_TYPE` | Cache type to use in service. Allowed values are `memory`, `redis`, `redis-cluster`. | `memory` | No |
| `CACHE_TTL` | Duration on how long to keep items in cache. Defaults to 0 meaning to never expire. | `0` | No |
| `CACHE_KEY_PREFIX` | Prefix all cache keys with specified value. | `""` | No |
| `CACHE_CONNECTION` | If other than memory cache is used specifies connection string on how to connect to cache storage. | `""` | No |
| `CACHE_PASSWORD` / `CACHE_PASSWORD_FILE` | Password to use in connection string. | `""` | No |
| `POSTGRES_HOST` | PostgreSQL HOST FQDN | `"db.example.lv"` | Yes |
| `POSTGRES_PORT` | PostgreSQL port | `"5432"` | Yes |
| `POSTGRES_USER` | PostgreSQL  | `"admin_public"` | Yes |
| `POSTGRES_DB` | PostgreSQL  | `"edim"` | Yes |
| `POSTGRES_PASSWORD` | PostgreSQL  | `/secret/edim-public-db-pw` | Yes |
| `IDAUTH_URL` | Authentication service URL | `""` | Yes |
| `IDAUTH_CLIENT_ID` | `api-person` id registrated in idAuth service | `""` | Yes |
| `IDAUTH_CLIENT_SECRET_FILE` | `api-person` secret registrated in idAuth service | `/secret/edim-idauth-client-secret-api-admin` | Yes |
| `AUDIT_ENDPOINT` | Audit endpoint | `"http://api-audit.edim-test.svc.cluster.local:8080/audit"` | Yes |

### Local example

In local development you must create `.env` file in the root of the project. Example:

```sh
    ENVIRONMENT: "Production"
    BASE_PATH: "/admin"
    LOG_LEVEL: "debug"
    REVERSE_PROXY_TRUSTED_IPS: "*"
    REVERSE_PROXY_LIMIT: "3"
    POSTGRES_HOST: "db.example.lv"
    POSTGRES_PORT: "5432"
    POSTGRES_USER: "admin_public"
    POSTGRES_DB: "edim"
    POSTGRES_PASSWORD_FILE: /secret/admin-public-db-pw
    IDAUTH_URL: ""
    IDAUTH_CLIENT_ID: ""
    IDAUTH_CLIENT_SECRET_FILE: /secret/edim-idauth-client-secret-api-admin
    AUDIT_ENDPOINT: "http://api-audit.edim-test.svc.cluster.local:8080/audit"
```
