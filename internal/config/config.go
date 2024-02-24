package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string     `yaml:"env" env-default:"local"`
	DB     Storage    `yaml:"storage"`
	Server HTTPServer `yaml:"http-server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:5000"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

type Storage struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
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
