package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/markelog/validate/logger"
)

// Up app
func Up() *iris.Application {
	var (
		app = iris.New()
		log = logger.Up()
	)

	app.Logger().Install(log)
	app.Use(recover.New())

	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 true,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

	return app
}
