package logging

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	cfg "github.com/elephant-insurance/go-utils/config"

	loggly "github.com/sebest/logrusly"
)

// ConstEmptyMessage => loggly throws a silent error if an empty message is provided
const ConstEmptyMessage = "."

// ConstLogFieldServerEvironment => the environment config value for the server running this code (ci, qa1, prod)
const ConstLogFieldServerEvironment = "env"

// ConstLogFieldServerHostName => the host name of the server running this code
const ConstLogFieldServerHostName = "host"

var isProd = checkIfProdEnv()

func checkIfProdEnv() bool {
	return strings.ToLower(cfg.Config.Environment) == "prod"
}

func init() {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// configure logging settings based on environment
	env := cfg.Config.Environment
	logLevel := log.DebugLevel // default log level = DEBUG

	if env == "dev" {
		// Log as TEXT and force colors in terminal
		log.SetFormatter(&log.TextFormatter{ForceColors: true})

		// Set log level = DEBUG
	} else {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		// Set log log level = INFO
		logLevel = log.InfoLevel
	}

	// configure logrus log level
	log.SetLevel(logLevel)

	// configure loggly
	hostName, _ := os.Hostname()
	logglyHook := loggly.NewLogglyHook(cfg.Config.LogglyKey, hostName, logLevel, env, "ms-customer-portal")
	log.Println("Adding Loggly logging hook")
	log.AddHook(logglyHook)

}

// LogError => logs a message and optional key/values with log level = ERROR
func LogError(errorMsg string, kvs *map[string]interface{}) {
	fields := log.Fields{}

	if kvs != nil {
		for key, value := range *kvs {
			fields[key] = value
		}
	}

	log.WithFields(fields).Error(errorMsg)
}
