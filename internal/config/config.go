package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage-path" env-required:"true"`
	Server      HTTPServer `yaml:"http-serever"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:5000"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG-PATH", "C:/Users/maus1/GolandProjects/pinterest-clone/config/local.yml")

	configPath := os.Getenv("CONFIG-PATH")

	if configPath == "" {
		log.Fatal("Config path is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	return &cfg
}
