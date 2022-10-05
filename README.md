# DRM ClearKey Server

[![Go](https://github.com/AlexandreBrg/go-clearkey-server/actions/workflows/go.yml/badge.svg)](https://github.com/AlexandreBrg/go-clearkey-server/actions/workflows/go.yml)

The idea behind this little project is to have a working ClearKey server for test purposes. This was the idea
behind the real first version, but I ended up with a new version of this project written with hexagonal 
architecture in mind. If you don't know what ClearKey systems are, I let you check what is a 
[DRM](https://www.fortinet.com/resources/cyberglossary/digital-rights-management-drm) first, then look for this kind 
of system. Of course, this server is completely compliant with [W3C](https://www.w3.org/TR/encrypted-media/#clear-key) 
specification.

As this is my first application in Go & with this kind of architecture, it is obvious that it is not production ready.

## Usage

```shell
go run .
```

If you are willing to use PostgreSQL as repository, you have a `docker-compose.yml` file available which have credentials
that are the same as environment variable's default.

### Video player

In order to try it, you can also run [index.html](./intelliJ/index.html), which contains a simple video player to run 
against your go-clearkey-server, please note that you need to update key ids accordingly. 

### Environment variables

| Variable name     | [Viper](https://github.com/spf13/viper) KeyID | Default                                                | Description                                                                                | 
|-------------------|-----------------------------------------------|--------------------------------------------------------|--------------------------------------------------------------------------------------------|
| `ENV`             | `EnvType`                                     | `development`                                          | Either `development` or `production`, if development the logs are sugared, else it is JSON |
| `PORT`            | `Port`                                        | `8080`                                                 | Listening port of the application                                                          | 
| `IP`              | `Ip`                                          | `0.0.0.0`                                              | Listening IP address of the application                                                    |
| `ALLOWED_DOMAINS` | `Domains`                                     | `[]string{"http://localhost:*", "http://127.0.0.1:*"}` | CORS allowed domains                                                                       |
| `REPOSITORY_TYPE` | `Repository`                                  | `RAM`                                                  | Define the type of repository used, choose between `RAM` and `PSQL`                        |
| `PSQL_PASSWORD`   | `Psql_pass`                                   | ` `                                                    | PostgreSQL password                                                                        |
| `PSQL_USER`       | `Psql_user`                                   | ` `                                                    | PostgreSQL username                                                                        |  
| `PSQL_ADDR`       | `Psql_addr`                                   | `127.0.0.0:5433`                                       | PostgreSQL address (default is `docker-compose.yml` port)                                  | 
| `PSQL_DB`         | `Psql_db`                                     | `postgres`                                             | PostgreSQL database                                                                        |
| `PSQL_INSECURE`   | `Psql_insecure`                               | `true`                                                 | Define whether postgres tries to connect using TLS handshake                               | 

## Routes

| Method | Route | Description |
| ------ | ----- | ----------- |
| POST | /license |  Request a license value following [W3C](https://www.w3.org/TR/encrypted-media/#clear-key-request-format) specs |
| POST | /license/register | Request to register a new license, no body expected, but it will generate a key that will be returned as [license format](https://www.w3.org/TR/encrypted-media/#clear-key-license-format) |

