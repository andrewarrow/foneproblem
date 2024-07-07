package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Invite(c *router.Context, second, third string) {
	if second != "" && third != "" && c.Method == "GET" {
		handleInvite(c, second, third)
		return
	}
	c.NotFound = true
}

func handleInvite(c *router.Context, guid, nrg string) {
	router.Redirect(c, "/core/workshops")
}
