package email

import (
	"github.com/kataras/iris/v12"
	"github.com/markelog/validate/modules/validate"
	"github.com/sirupsen/logrus"
)

// Email is an object for route params to validate
type Email struct {
	Email string `json:"email,omitempty"`
}

// Up project route
func Up(app *iris.Application, log *logrus.Logger) error {
	app.Post("/email/validate", func(ctx iris.Context) {
		var params Email
		ctx.ReadJSON(&params)

		valid, validators := validate.New(params.Email).Validate("email")

		log.WithFields(logrus.Fields{
			"email": params.Email,
		}).Info("Email checked")

		if valid {
			ctx.StatusCode(iris.StatusOK)
		} else {
			ctx.StatusCode(iris.StatusBadRequest)
		}

		ctx.JSON(iris.Map{
			"valid":      valid,
			"validators": validators,
		})
	})

	return nil
}
