package common_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/common"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
)

var (
	app *iris.Application

	// Response schema
	response = `{
		"type": "object",
	    "properties": {
			"message": {"type": "string"},
			"status":  {"type": "string"}
	    },
	    "required": ["message", "status"]
	}`
)

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	common.Up(app, log)

	app.Build()

	os.Exit(m.Run())
}

func TestPOST404(t *testing.T) {
	req := request.Up(app, t)

	response := req.POST("/not-found").
		WithHeader("Content-Type", "routes/json").
		Expect().
		Status(http.StatusNotFound)

	json := response.JSON()

	json.Schema(response)

	json.Object().Value("status").Equal("error")
	json.Object().Value("message").Equal("Can't find this route")
}
