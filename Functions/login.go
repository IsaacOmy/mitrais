package functions

import (
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {

	Tmp.ExecuteTemplate(res, "login", nil)
}
