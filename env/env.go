package env

import (
	"log"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var envOnce sync.Once
var envSingleton Envars

// Envars contains all the environment variables for the application
type Envars struct {
	AppName         string        `split_words:"true" default:"array"`
	AppDomain       string        `split_words:"true" default:"array.com"`
	Port            string        `split_words:"true" default:":8080"`
	DatabaseURL     string        `split_words:"true" required:"true"`
	TablePrefix     string        `split_words:"true" default:""`
	RedisAddr       string        `split_words:"true" default:"localhost:6379"`
	SessionLifetime time.Duration `split_words:"true" default:"24h"`
}

// GetEnvironment returns a singleton object contains all environment variables
func GetEnvironment() *Envars {
	envOnce.Do(func() {
		if err := envconfig.Process("", &envSingleton); err != nil {
			log.Fatalln(err)
		}

		// required to parse time.Time values
		envSingleton.DatabaseURL += "?parseTime=True"
	})

	return &envSingleton
}
