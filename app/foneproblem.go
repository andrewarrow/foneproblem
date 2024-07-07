package app

import (
	"foneproblem/models"

	"github.com/andrewarrow/feedback/router"
)

func FoneProblem(c *router.Context, second, third string) {
	if second == "start" && third == "" && c.Method == "GET" {
		handleFoneProblemIndex(c)
		return
	}
	if second == "workshops" && third == "" && c.Method == "GET" {
		handleWorkshopsIndex(c)
		return
	}
	if second == "dashboard" && third == "" && c.Method == "GET" {
		handleDashboard(c)
		return
	}
	if second == "logout" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	c.NotFound = true
}

func handleWorkshopsIndex(c *router.Context) {

	c.Title = "workshops in your area | foneproblem.com"
	one := c.One("event", "order by created_at desc")
	send := map[string]any{}
	send["guid"] = one["guid"]

	eMap := models.EnergyMap()
	items := c.All("event_member", "where event_id=$1 order by nrg", "", one["id"])
	for _, item := range items {
		nrg, _ := item["nrg"].(string)
		eMap[nrg] = false
	}
	send["items"] = items
	send["emap"] = eMap
	send["nrgs"] = models.AllEnergiesHuman()

	c.SendContentInLayout("workshops.html", send, 200)
}

func handleFoneProblemIndex(c *router.Context) {

	c.Title = "do you have a foneproblem? | foneproblem.com"
	send := map[string]any{}
	send["questions"] = questions
	c.SendContentInLayout("start.html", send, 200)
}

var questions = []string{"Do you often find yourself using your smartphone longer than you intended?", "Do you feel anxious or irritable when you do not have access to your smartphone?", "Do you frequently check your smartphone even when it hasn't notified you of anything?", "Have you ever tried to cut down on your smartphone use and failed?", "Do you find that you use your smartphone in inappropriate or dangerous situations, such as while driving?", "Does your smartphone use interfere with your daily activities, work, or relationships?", "Do you use your smartphone to escape from stress or negative feelings?", "Do you often lose track of time because you are using your smartphone?", "Have your friends or family members expressed concern about the amount of time you spend on your smartphone?"}
