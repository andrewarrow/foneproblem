package app

import "github.com/andrewarrow/feedback/router"

func FoneProblem(c *router.Context, second, third string) {
	if second == "start" && third == "" && c.Method == "GET" {
		handleFoneProblemIndex(c)
		return
	}
	c.NotFound = true
}

func handleFoneProblemIndex(c *router.Context) {

	c.Title = "do you have a foneproblem? | foneproblem.com"
	send := map[string]any{}
	c.SendContentInLayout("start.html", send, 200)
}
