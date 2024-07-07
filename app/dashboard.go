package app

import (
	"fmt"
	"foneproblem/models"

	"github.com/andrewarrow/feedback/router"
)

func handleDashboard(c *router.Context) {
	c.Title = "dashboard | foneproblem.com"
	sql := fmt.Sprintf("where id in (select event_id from event_members where user_id=%d) order by created_at desc", c.User["id"])
	items := c.All("event", sql, "")

	for _, item := range items {
		eMap := models.EnergyMap()
		userMap := map[string]any{}
		subitems := c.All("event_member", "where event_id=$1", "", item["id"])
		c.DecorateList(subitems)
		for _, sub := range subitems {
			nrg, _ := sub["nrg"].(string)
			eMap[nrg] = false
			userMap[nrg] = sub["user"]
			sub["nrg_human"] = models.Energy(sub["nrg"])
		}
		item["members"] = subitems
		item["emap"] = eMap
		item["usermap"] = userMap
	}
	send := map[string]any{"items": items, "nrgs": models.AllEnergiesHuman()}
	c.SendContentInLayout("dashboard.html", send, 200)
}
