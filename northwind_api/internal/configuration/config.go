package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	get(key string) string
}
type config struct {
}

func New() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &config{}
}
func (c *config) get(key string) string {
	return os.Getenv(key)
}
