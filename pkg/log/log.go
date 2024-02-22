package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	file, err := os.OpenFile("cdn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to log to file, using default stderr")
	}

	log.SetOutput(file)
}

func LogInfo(msg string) {
	log.Info(msg)
}

func LogTrace(msg string) {
	log.Trace(msg)
}

func LogError(msg string) {
	log.Error(msg)
}

func LogFatal(msg string) {
	log.Fatal(msg)
}
