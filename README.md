# URL Shortener
The online service that can transfrom long url into short url.

# Require
* MySQL
* Redis

# Install
```

go build ./main.go

// before you run it you must make sure the service mysql and redis service had been running
// configure ./conf/server.config.json

./main.exe


```

# Configuration
```
{
    "server": "localhost",  // eg. http://pwcong.me which decide the short url prefix that can generate short url like http://pwcong.me/ABC
    "host": "localhost",    // server local ip address
    "port": "80",           // server listening port
    "db": {
        "redis": {
            "address": "localhost:6379" // redis service listening address
        },
        "mysql": {
            "address": "localhost:3306",    // mysql service listening address
            "dbname":"url_shortener",       // database name
            "user":"pwcong",                
            "password":"123456"
        }
    }
}
```