# Bouncer - API to manage users and authentication

This project has to main goal to provide an easy backend that manage users and authentication.
It provides authentication by linking users to a [JWT](http://jwt.io) token.

## Status

Bouncer is currently under development. Feel free to comment or contribute!

## Building

Bouncer is pretty simple to build.

- Clone the repository
- Install `glide` if not already present
- Install dependencies with glide (ie `glide install`)
- Build and run the application (ie `go run bouncer.go`)

## Usage

```
$ go run bouncer.go --help

Expose authentication and users

Usage:
  bouncer [flags]
  bouncer [command]

Available Commands:
  help        Help about any command
  version     Print the version number

Flags:
      --config string              Set configuration file
  -h, --help                       help for bouncer
      --jwt-realm string           Set the JWT application realm (default "localhost")
      --jwt-secret string          Set the JWT application secret
      --port uint                  Set port to listen (default 8080)
      --postgres-dbname string     Set the name of the database to connect (default "bouncer")
      --postgres-host string       Set ip address of postgres (default "127.0.0.1")
      --postgres-password string   Set password of the database
      --postgres-sslmode string    Set the sslmode of the postgres client (default "disable")
      --postgres-user string       Set user of the database (default "postgres")
      --verbose                    Set output to verbose

```

## Routes

A list of available routes:

|Path|Method|Description|
|---|---|---|
|`/api/auth`|`POST`|Get a JWT you need to give `username` and `password` as a JSON body|
|`/api/user`|`GET`|Get information about all users|
|`/api/user`|`POST`|Create an user|
|`/api/user/:uuid`|`GET`|Get information about one user identify by `uuid`|
|`/api/user/:uuid`|`PUT`|Update information about one user identify by `uuid`|
|`/api/user/:uuid`|`DELETE`|Soft delete information about one user identify by `uuid`|

### Details

Following details about routes and body content that consumes

- `/api/auth`(`POST`): Get a JWT you need to give `username` and `password` as a JSON body

No particulary header to set.

```json
{
  "username": "Jane",
  "password": "Doe"
}
```

- `/api/user`(`GET`): Get information about all users

You need to set the header `Authorization` as this `Authorization: JWT <json-web-token>` with the token return by the route `/api/auth`(`POST`)

No request body is needed

- `/api/user`(`POST`): Create an user

You need to set the header `Authorization` as this `Authorization: JWT <json-web-token>` with the token return by the route `/api/auth`(`POST`)

```json
{
  "FirstName": "Jane",
  "LastName": "Doe",
  "Email": "jane.doe@anonym.ous",
  "Password": "§4[_(_)†"
}
```

- `/api/user/:uuid`(`GET`): Get information about one user identify by `uuid`

You need to set the header `Authorization` as this `Authorization: JWT <json-web-token>` with the token return by the route `/api/auth`(`POST`)

No request body is needed

- `/api/user/:uuid`(`PUT`): Update information about one user identify by `uuid`

Information to update can be only one or more field of the following JSON. Only information which is given in the request body will be updated.

```json
{
  "ID": "<uuid>",
  "FirstName": "Jane",
  "LastName": "Doe",
  "Email": "jane.doe@anonym.ous",
  "Password": "§4[_(_)†",
  "CreatedAt": "2017-06-13T14:21:09.256165Z",
  "UpdatedAt": "2017-06-13T14:21:09.256165Z",
  "DeletedAt": null
}
```

You need to set the header `Authorization` as this `Authorization: JWT <json-web-token>` with the token return by the route `/api/auth`(`POST`)

- `/api/user/:uuid`(`DELETE`): Soft delete information about one user identify by `uuid`

You need to set the header `Authorization` as this `Authorization: JWT <json-web-token>` with the token return by the route `/api/auth`(`POST`)

No request body is needed

## Configuration

Bouncer can read a configuration file.

Configuration is load and override in the following order:

- /etc/bouncer/config.yaml
- ~/bouncer/config.yaml
- ./config.yaml
- config filepath from command line
