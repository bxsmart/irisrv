package public

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

func InitYaag(app *iris.Application) {
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "views/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})

	app.Use(irisyaag.New())
}
