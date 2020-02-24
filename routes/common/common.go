package common

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

// Up project route
func Up(app *iris.Application, log *logrus.Logger) {
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		log.WithFields(logrus.Fields{
			"url":    ctx.Path(),
			"method": ctx.Method(),
		}).Error("Returned 404 error")

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"status":  "error",
			"message": "Can't find this route",
		})
	})
}
