# URL Shortener
The online service that can transfrom long url into short url.

# Require
* MySQL
* Redis

# Install & Usage
```

go build ./main.go

// before you run it you must make sure the service mysql and redis had been running
// configure ./conf/server.config.json

./main.exe


```

**URL Shortener Service API:**
${server|host}:${port}/api?url=${long_url}[&format=json]

```
eg.

localhost/api?url=http://www.pwcong.me/     
# localhost/ABCD

localhost/api?url=http://www.pwcong.me/&format=json     
# {"Err":"","LongUrl":"http://www.pwcong.me" "ShortUrl":"localhost/ABCD"}

http://pwcong.me/api?url=https://github.com/
# http://pwcong.me/EFGH

127.0.0.1:8080/api?url=http://www.pwcong.me/
# localhost:8080/ABCD
```

# Configuration
```
{
    "mode": "prod",         // prod | dev
    "domain": "localhost",  // eg. pwcong.me which decide the short url prefix that can generate short url like http://pwcong.me/ABC
    "host": "0.0.0.0",      // server public ip address
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