package app

import "github.com/andrewarrow/feedback/router"

func handleDashboard(c *router.Context) {
	c.Title = "dashboard | foneproblem.com"
	send := map[string]any{}
	c.SendContentInLayout("dashboard.html", send, 200)
}
