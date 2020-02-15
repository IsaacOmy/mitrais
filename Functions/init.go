package functions

import (
	"html/template"
)

var Tmp = template.Must(template.ParseFiles(
	"View/header.html",
	"View/footer.html",
	"View/registration.html",
	"View/login.html",
))
