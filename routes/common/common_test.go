package common_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/markelog/validate/logger"
	"github.com/markelog/validate/routes/common"
	"github.com/markelog/validate/test/request"
	"github.com/markelog/validate/test/routes"
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
	app = routes.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	common.Up(app, log)

	app.Build()

	os.Exit(m.Run())
}

func TestPOST404(t *testing.T) {
	req := request.Up(app, t)

	common := req.POST("/not-found").
		WithHeader("Content-Type", "routes/json").
		Expect().
		Status(http.StatusNotFound)

	json := common.JSON()

	json.Schema(response)

	json.Object().Value("status").Equal("error")
	json.Object().Value("message").Equal("Can't find this route")
}
