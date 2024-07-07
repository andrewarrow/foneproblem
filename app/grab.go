package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Grab(c *router.Context, second, third string) {
	if second != "" && third != "" && c.Method == "GET" {
		handleWorkshopGrab(c, second, third)
		return
	}
	c.NotFound = true
}

func handleWorkshopGrab(c *router.Context, guid, nrg string) {

	router.Redirect(c, "/core/dashboard")
}
