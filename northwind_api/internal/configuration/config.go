package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}
type config struct{}

func NewConfig(files ...string) Config {
	err := godotenv.Load(files...)
	if err != nil {
		panic(err)
	}
	return &config{}
}
func (c *config) Get(key string) string {
	return os.Getenv(key)
}
