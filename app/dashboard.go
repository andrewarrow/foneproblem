package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func handleDashboard(c *router.Context) {
	c.Title = "dashboard | foneproblem.com"
	sql := fmt.Sprintf("where id in (select event_id from event_members where user_id=%d) order by created_at desc", c.User["id"])
	items := c.All("event", sql, "")
	send := map[string]any{"items": items}
	c.SendContentInLayout("dashboard.html", send, 200)
}
