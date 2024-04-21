package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var Cfg *Config

type Config struct {
	Port     string   `yaml:"port"`
	Log      Log      `yaml:"log"`
	Database Database `yaml:"database"`
}

type Log struct {
	Path string `yaml:"path"`
}

type Database struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	DbName string `yaml:"dbname"`
}

// init the config into global config point from the config file
func init() {
	cfgFile := flag.String("config_file", "", "config file path")
	flag.Parse()

	if "" != *cfgFile || len(strings.TrimSpace(*cfgFile)) != 0 {
		cfgContent, err := os.ReadFile(*cfgFile)
		if err != nil {
			err := fmt.Errorf("read config file error: %s", err.Error())
			panic(err)
		}
		Cfg = &Config{}
		err = yaml.Unmarshal(cfgContent, Cfg)
		if err != nil {
			err := fmt.Errorf("unmarshal config file error: %s", err.Error())
			panic(err)
		}
	}
}

// init the config into global config point from the env
func init() {
	port, portExist := os.LookupEnv("HTTP_PORT")
	host, hostExist := os.LookupEnv("DB_HOST")
	user, userExist := os.LookupEnv("DB_USER")
	passwd, passwdExist := os.LookupEnv("DB_PASSWORD")
	dbName, dbNameExist := os.LookupEnv("DB_NAME")
	if portExist {
		Cfg.Port = port
	}
	if hostExist {
		Cfg.Database.Host = host
	}
	if userExist {
		Cfg.Database.User = user
	}
	if passwdExist {
		Cfg.Database.Passwd = passwd
	}
	if dbNameExist {
		Cfg.Database.DbName = dbName
	}
}
