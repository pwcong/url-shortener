# URL Shortener
The online service that can transfrom long url into short url.

## Require
* MySQL
* Redis

## Install & Usage
```shell
> go get .
> go build

// before you run it you must make sure the service mysql and redis had been running
// configure ./conf/default.toml

> ./url-shortener.exe

```

## Configuration

```toml
[server]
host = "0.0.0.0"
port = 7001

[middlewares]

    [middlewares.cors]
    active = true
    [middlewares.logger]
    active = true

[databases]

    [databases.mysql]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "root"
    dbname = "url_shortener"

    [databases.redis]
    host = "127.0.0.1"
    port = 6379

```