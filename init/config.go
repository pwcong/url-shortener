package init

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const DEFAULT_CONFIG = `{
    "server": "localhost",
    "host": "localhost",
    "port": "80",
    "db": {
        "redis": {
            "address": "localhost:6379"
        },
        "mysql": {
            "address": "localhost:3306",
            "dbname":"url_shortener",
            "user":"pwcong",
            "password":"123456"
        }
    }
}`

type MySQLConfig struct {
	Address  string
	DBName   string
	User     string
	Password string
}

type RedisConfig struct {
	Address string
}

type config struct {
	Server string
	Host   string
	Port   string
	DB     struct {
		Redis RedisConfig
		MySQL MySQLConfig
	}
}

var Config config

func loadConfig() {

	var config interface{}

	configPath := filepath.Join(filepath.Dir(os.Args[0]), "conf", "server.config.json")

	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		json.Unmarshal([]byte(DEFAULT_CONFIG), &config)
	} else {
		json.Unmarshal(data, &config)
	}

	configMap, ok := config.(map[string]interface{})
	if !ok {
		log.Fatal("can not load config")
	}

	///////////////////////////////////////////////////////
	/*********** load configuration of addr **************/
	///////////////////////////////////////////////////////
	configHostIF, ok := configMap["host"]
	if !ok {
		log.Fatal("can not load config.host")
	}

	configHost, ok := configHostIF.(string)
	if !ok {
		log.Fatal("can not load config.host")
	}

	Config.Host = configHost

	configServerIF, ok := configMap["server"]
	if !ok {
		log.Fatal("can not load config.server")
	}

	configServer, ok := configServerIF.(string)
	if !ok {
		log.Fatal("can not load config.server")
	}

	Config.Server = configServer

	configPortIF, ok := configMap["port"]
	if !ok {
		log.Fatal("can not load config.server")
	}

	configPort, ok := configPortIF.(string)
	if !ok {
		log.Fatal("can not load config.server")
	}

	Config.Port = configPort

	/////////////////////////////////////////////////////
	/*********** load configuration of db **************/
	/////////////////////////////////////////////////////
	configDBIF, ok := configMap["db"]
	if !ok {
		log.Fatal("can not load config.db")
	}

	configDBMap, ok := configDBIF.(map[string]interface{})
	if !ok {
		log.Fatal("can not load config.db")
	}

	////////////////////////////////////////////////////////
	/*********** load configuration of mysql **************/
	////////////////////////////////////////////////////////
	configDBMySQLIF, ok := configDBMap["mysql"]
	if !ok {
		log.Fatal("can not load config.db.mysql")
	}

	configDBMySQLMap, ok := configDBMySQLIF.(map[string]interface{})
	if !ok {
		log.Fatal("can not load config.db.mysql")
	}

	configDBMySQLAddressIF, ok := configDBMySQLMap["address"]
	if !ok {
		log.Fatal("can not laod config.db.mysql.address")
	}

	configDBMySQLAddress, ok := configDBMySQLAddressIF.(string)
	if !ok {
		log.Fatal("can not laod config.db.mysql.address")
	}

	Config.DB.MySQL.Address = configDBMySQLAddress

	configDBMySQLDBNameIF, ok := configDBMySQLMap["dbname"]
	if !ok {
		log.Fatal("can not laod config.db.mysql.dbname")
	}

	configDBMySQLDBName, ok := configDBMySQLDBNameIF.(string)
	if !ok {
		log.Fatal("can not laod config.db.mysql.dbname")
	}

	Config.DB.MySQL.DBName = configDBMySQLDBName

	configDBMySQLUserIF, ok := configDBMySQLMap["user"]
	if !ok {
		log.Fatal("can not laod config.db.mysql.user")
	}

	configDBMySQLUser, ok := configDBMySQLUserIF.(string)
	if !ok {
		log.Fatal("can not laod config.db.mysql.user")
	}

	Config.DB.MySQL.User = configDBMySQLUser

	configDBMySQLPasswordIF, ok := configDBMySQLMap["password"]
	if !ok {
		log.Fatal("can not laod config.db.mysql.password")
	}

	configDBMySQLPassword, ok := configDBMySQLPasswordIF.(string)
	if !ok {
		log.Fatal("can not laod config.db.mysql.password")
	}

	Config.DB.MySQL.Password = configDBMySQLPassword

	////////////////////////////////////////////////////////
	/*********** load configuration of redis **************/
	////////////////////////////////////////////////////////

	configDBRedisIF, ok := configDBMap["redis"]
	if !ok {
		log.Fatal("can not load config.db.redis")
	}

	configDBRedisMap, ok := configDBRedisIF.(map[string]interface{})
	if !ok {
		log.Fatal("can not load config.db.redis")
	}

	configDBRedisAddressIF, ok := configDBRedisMap["address"]
	if !ok {
		log.Fatal("can not load config.db.redis.address")
	}

	configDBRedisAddress, ok := configDBRedisAddressIF.(string)
	if !ok {
		log.Fatal("can not load config.db.redis.address")
	}

	Config.DB.Redis.Address = configDBRedisAddress
}

func init() {

	now := time.Now().Unix()
	loadConfig()
	log.Printf("configuration has been loaded in %v", time.Now().Unix()-now)

}
