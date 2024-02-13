package config

import (
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) IConfig {
	if err := godotenv.Load(filenames...); err != nil {
		panic("error load env")
	}
	return &configImpl{}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
