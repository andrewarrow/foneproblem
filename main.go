package main

import (
	"embed"
	"foneproblem/app"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]

	if arg == "seed" {
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		c := r.ToContext()
		for _, city := range app.MajorCities {
			c.Params = map[string]any{}
			c.Params["location"] = city
			c.ValidateAndInsert("event")
		}
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.FoneProblem
		r.Paths["register"] = app.Register
		r.Paths["grab"] = app.Grab
		r.Paths["invite"] = app.Invite
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}
}
