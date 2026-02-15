package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	//keyEnv := os.Getenv("KEY")
	//fmt.Println("keyEnv:", keyEnv)
	return &Config{Key: ""}
}

func (config *Config) GetKey() {
	keyEnv := os.Getenv("KEY")
	config.Key = keyEnv
	fmt.Println("keyEnv:", keyEnv)
}
