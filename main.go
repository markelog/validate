package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/markelog/validate/env"
	"github.com/markelog/validate/logger"
	"github.com/markelog/validate/routes"
	"github.com/markelog/validate/routes/email"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Up()

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	var (
		app = routes.Up()
		log = logger.Up()
	)

	if err := email.Up(app, log); err != nil {
		log.Panic(err)
	}

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
