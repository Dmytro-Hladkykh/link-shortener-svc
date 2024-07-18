# link-shortener-svc

## Description

This service provides URL shortening functionality. It allows users to create short links from long URLs, redirect users to original URLs when they use shortened links.

## Install

```
git clone github.com/Dmytro-Hladkykh/link-shortener-svc
cd link-shortener-svc
go build main.go
export KV_VIPER_FILE=./config.yaml
./main migrate up
./main run service
```

## Running from Docker

Make sure that docker installed.

Configure `docker-compose.yml` with entrypoint commands:

```
entrypoint:
      [
        "/bin/sh",
        "-c",
        "/usr/local/bin/link-shortener-svc migrate up && /usr/local/bin/link-shortener-svc run service",
      ]
```

Possible commands are:

```
/usr/local/bin/link-shortener-svc migrate up
```

```
/usr/local/bin/link-shortener-svc run service
```

```
/usr/local/bin/link-shortener-svc migrate down
```

To run the service with Docker use:

```
docker-compose up --build
```

## Testing

To test usage you can use Postman.

### Create a POST with:

```
http://localhost:8000/link-shortener
```

With Body:

```
{
  "original_url": "http://example.com"
}
```

In response you will get:

```
{
  "short_code": "YrEiin"
}
```

### Create a GET with:

```
http://localhost:8000/link-shortener/YrEiin
```

In response you will get:

```
{
  "original_url": "http://example.com"
}
```

## Running from Source

- Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
- Provide valid config file
- Launch the service with `migrate up` command to create database schema
- Launch the service with `run service` command

### Database

For services, we do use **_PostgresSQL_** database.
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).

### Third-party services

## Contact

Responsible Dmytro Hladkykh
The primary contact for this project is https://t.me/Dimo4kaaaaaa
