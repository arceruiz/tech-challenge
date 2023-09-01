package config

import (
	"flag"
	"log"
	"os"

	"github.com/notnull-co/cfg"
)

var (
	Cfg Config
)

type Config struct {
	Token struct {
		Key string `cfg:"key"`
	} `cfg:"token"`
	Server struct {
		Port string `cfg:"port"`
	} `cfg:"server"`
	DB struct {
		ConnectionString string `cfg:"connectionString"`
	} `cfg:"db"`
}

func ParseFromFlags() {
	os.Setenv("APP_DB_CONNECTIONSTRING", "host=172.19.0.1 port=5432 dbname=fiap_tech_challenge user=postgres password=1234 sslmode=disable")
	os.Setenv("APP_TOKEN_KEY", "vuIXaOK4OpJWA9ySX1UTpIWshXPpP6neGKGA724FauY")
	var configDir string

	flag.StringVar(&configDir, "config-dir", "../../internal/config/", "Configuration file directory")
	flag.Parse()

	parse(configDir)
}

func parse(dirs ...string) {
	if err := cfg.Load(&Cfg,
		cfg.Dirs(dirs...),
		cfg.UseEnv("app"),
	); err != nil {
		log.Panic(err)
	}
}
