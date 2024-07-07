package app

import "github.com/andrewarrow/feedback/router"

func Register(c *router.Context, second, third string) {
	if second != "" && third != "" && c.Method == "GET" {
		handleWorkshopRegister(c, second, third)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}
	c.NotFound = true
}

func handleWorkshopRegister(c *router.Context, guid, nrg string) {

	event := c.One("event", "where guid=$1", guid)
	c.FreeFormUpdate("insert into event_members (user_id, event_id, nrg) values ($1, $2, $3)", c.User["id"], event["id"], nrg)

	c.Title = "register | foneproblem.com"
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
