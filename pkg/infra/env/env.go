package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorln("Error loading .env file, using default environment variables")
	}
}
