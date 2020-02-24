package env

import (
	"github.com/joho/godotenv"
	"github.com/markelog/validate/logger"
)

// Up environment
func Up() {
	log := logger.Up()

	err := godotenv.Load()
	if err != nil {
		log.Info("Haven't load the .env file")
	}
}
