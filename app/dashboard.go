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

	eMap := models.EnergyMap()

	for _, item := range items {
		subitems := c.All("event_member", "where event_id=$1", "", item["id"])
		for _, sub := range subitems {
			nrg, _ := sub["nrg"].(string)
			eMap[nrg] = false
			sub["nrg_human"] = models.Energy(sub["nrg"])
		}
		c.DecorateList(subitems)
		item["members"] = subitems
	}
	send := map[string]any{"items": items, "nrgs": models.AllEnergiesHuman()}
	c.SendContentInLayout("dashboard.html", send, 200)
}
