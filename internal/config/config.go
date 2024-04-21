package config

var Cfg *Config

type Config struct {
	Port     string   `yaml:"port"`
	Log      Log      `yaml:"log"`
	Database Database `yaml:"database"`
}

type Log struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type Database struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	DbName string `yaml:"dbname"`
}
