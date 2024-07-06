package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/foneproblem/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "{{/homeducky}}/start")
	}
	if Global.Start == "start.html" {
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "foneproblem", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "foneproblem", nil, afterRegister)
	}
}
