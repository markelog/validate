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

		validator := validate.New(params.Email)
		valid, validators := validator.Validate("email")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"valid":      valid,
			"validators": validators,
		})
	})

	return nil
}
