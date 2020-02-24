package email_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/markelog/validate/env"
	"github.com/markelog/validate/logger"
	"github.com/markelog/validate/routes"
	"github.com/markelog/validate/routes/email"
	"github.com/markelog/validate/test/request"
)

var (
	app *iris.Application

	// response schema
	response = `{
		"type": "object",
	    "properties": {
			"valid": {"type": "boolean"},
			"validators": {"type": "object"}
	    },
	    "required": ["valid", "validators"]
	}`
)

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	email.Up(app, log)

	app.Build()

	os.Exit(m.Run())
}

func TestSuccess(t *testing.T) {
	req := request.Up(app, t)
	data := map[string]interface{}{
		"email": "markelog@gmail.com",
	}

	validate := req.POST("/email/validate").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	json := validate.JSON()
	validators := json.Object().Value("validators").Object()

	json.Schema(response)

	json.Object().Value("valid").Equal(true)
	validators.Value("regexp").Object().Value("valid").Equal(true)
	validators.Value("dmarc").Object().Value("valid").Equal(true)
	validators.Value("domain").Object().Value("valid").Equal(true)
	validators.Value("smtp").Object().Value("valid").Equal(true)

}

func TestBadEmail(t *testing.T) {
	req := request.Up(app, t)
	data := map[string]interface{}{
		"email": "nope",
	}

	validate := req.POST("/email/validate").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := validate.JSON()
	validators := json.Object().Value("validators").Object()

	json.Schema(response)

	json.Object().Value("valid").Equal(false)
	validators.Value("regexp").Object().Value("valid").Equal(false)
	validators.Value("dmarc").Object().Value("valid").Equal(false)
	validators.Value("domain").Object().Value("valid").Equal(false)
	validators.Value("smtp").Object().Value("valid").Equal(false)
}
