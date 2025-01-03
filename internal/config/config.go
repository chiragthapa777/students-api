package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// if function has Most prefix then the general practice is not to return error, is error the stop the program
func MustLoad() *Config {
	var configPath string = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// from args while running the programs eg go run main.go -config ./config.yaml
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is no set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}
