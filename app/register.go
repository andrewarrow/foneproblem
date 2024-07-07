package app

import (
	"fmt"
	"math/rand"

	"github.com/andrewarrow/feedback/router"
)

func Register(c *router.Context, second, third string) {
	if second != "" && third != "" && c.Method == "GET" {
		handleWorkshopRegister(c, second, third)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		userGuid := router.HandleCreateUserAutoForm(c, "")
		event_guid := router.GetCookie(c, "event_guid")
		nrg := router.GetCookie(c, "nrg")
		cookieToken := c.One("cookie_token", "where guid=$1", userGuid)

		event := c.One("event", "where guid=$1", event_guid)
		c.FreeFormUpdate("insert into event_members (user_id, event_id, nrg) values ($1, $2, $3)", cookieToken["user_id"], event["id"], nrg)

		return
	}
	c.NotFound = true
}

func handleWorkshopRegister(c *router.Context, guid, nrg string) {

	c.Title = "register | foneproblem.com"
	router.SetCookie(c, "event_guid", guid)
	router.SetCookie(c, "nrg", nrg)
	send := map[string]any{}
	send["guid"] = guid
	send["nrg"] = nrg
	send["email"] = fmt.Sprintf("oneone+test%006d@gmail.com", rand.Intn(999999))
	c.SendContentInLayout("register.html", send, 200)
}
