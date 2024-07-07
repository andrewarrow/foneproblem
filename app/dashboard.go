package app

import "github.com/andrewarrow/feedback/router"

func handleDashboard(c *router.Context) {
	c.Title = "dashboard | foneproblem.com"
	items := c.All("event", "where id in (select user_id from event_members where user_id=$1) order by created_at desc", "")
	send := map[string]any{"items": items}
	c.SendContentInLayout("dashboard.html", send, 200)
}
