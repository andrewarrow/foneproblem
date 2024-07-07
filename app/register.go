package app

import "github.com/andrewarrow/feedback/router"

func Register(c *router.Context, second, third string) {
	if second != "" && third != "" && c.Method == "GET" {
		handleWorkshopRegister(c, second, third)
		return
	}
	c.NotFound = true
}

func handleWorkshopRegister(c *router.Context, guid, nrg string) {

	c.Title = "register | foneproblem.com"
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
