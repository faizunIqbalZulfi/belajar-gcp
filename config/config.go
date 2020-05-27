package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Conf Config

func init() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	ConfigFile, err := ioutil.ReadFile("./config.json")
	switch {
	case os.IsNotExist(err):
		ReadEnvConfig(&Conf)
	case err == nil:
		if err := json.Unmarshal(ConfigFile, &Conf); err != nil {
			log.Println(" ====== ", os.IsNotExist(err))
			log.Println(" ====== ", err.Error())
			ReadEnvConfig(&Conf)
		}
	}
}

type Config struct {
	Db         PSQL             `json:"psql"`
	HTTPServer HTTPServerConfig `json:"httpServer"`
}

type HTTPServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type PSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

func (p PSQL) ConnectionString() string {
	return p.ConnectionStringWithSSLMode("disable")
}

func (p PSQL) ConnectionStringWithSSLMode(sslmode string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		p.Host, p.User, p.Password, p.DBName, p.Schema, sslmode)

}

func ReadEnvConfig(c *Config) {
	c.Db.Host = os.Getenv("PSQL_SERVER_HOST")
	c.Db.Port = os.Getenv("PSQL_SERVER_PORT")
	c.Db.DBName = os.Getenv("PSQL_SERVER_NAME")
	c.Db.User = os.Getenv("PSQL_SERVER_USERNAME")
	c.Db.Password = os.Getenv("PSQL_SERVER_PASSWORD")
	c.Db.Schema = os.Getenv("PSQL_SERVER_SCHEMA")

	c.HTTPServer.Host = os.Getenv("GOLOAN_SERVICE_HOST")
	c.HTTPServer.Port = os.Getenv("GOLOAN_SERVICE_PORT")

}
