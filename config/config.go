package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config is used as a global reference to the configuration
var Config *config

func init() {
	Config = loadConfiguration()
}

// Configuration struct to hold config info
type config struct {
	Environment string
	LogglyKey   string
}

// LoadConfiguration loads and returns a pointer to the loaded config file
func loadConfiguration() *config {
	c := &config{}

	// get the current environment
	c.Environment = os.Getenv("CONFIG_ENVIRONMENT")

	if c.Environment == "" {
		fmt.Println("unable to locate 'env' environment variable")
	}

	viper.SetConfigName("config")
	wd, err1 := os.Getwd()
	viper.AddConfigPath(wd)

	if err1 != nil {
		fmt.Println("No configuration file loaded")
	}

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded")
	}

	c.LogglyKey = viper.GetString(c.Environment + ".logglyKey")

	return c
}

// GetConfigValue returns the value of the key in the configuration file
func GetConfigValue(key string) string {
	return viper.GetString(Config.Environment + key)
}

// GetConfigValues returns a list of values based on the key in the configuration file
func GetConfigValues(key string) []string {
	return viper.GetStringSlice(Config.Environment + key)
}

// GetEnvironmentValue gets the value of an environment variable
func GetEnvironmentValue(key string) string {
	return os.Getenv(key)
}
