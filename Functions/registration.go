package functions

import (
	"encoding/json"
	"net/http"

	"github.com/IsaacOmy/mitrais/db"

	"github.com/IsaacOmy/mitrais/Model"
)

func Registration(res http.ResponseWriter, req *http.Request) {

	Tmp.ExecuteTemplate(res, "registration", nil)
}

func SubmitRegistration(res http.ResponseWriter, req *http.Request) {
	users := Model.Users{}
	successResponse := Model.SuccessResponse{}

	users.FirstName = req.FormValue("first_name")
	users.LastName = req.FormValue("last_name")
	users.Email = req.FormValue("email")
	users.Gender = req.FormValue("gender")
	users.MobileNumber = req.FormValue("mobile_number")

	dt := req.FormValue("date")
	if len(dt) <= 1 {
		dt = "0" + dt
	}
	mt := req.FormValue("month")
	yr := req.FormValue("year")
	users.DateOfBirth = yr + "-" + mt + "-" + dt

	errResponse := RegistrationValidate(users)
	res.Header().Set("Content-Type", "application/json")
	if len(errResponse.Error) > 0 {
		js, err := json.Marshal(errResponse)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(js)
	} else {
		db := db.GetDB()
		db.Create(&users)
		successResponse.Status = "data save"
		js, err := json.Marshal(successResponse)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(js)
	}
}

func RegistrationValidate(usr Model.Users) Model.ErrorResponse {
	errors := Model.ErrorResponse{}
	db := db.GetDB()
	emailValidation := Model.Users{}
	if !db.Where("email = ?", usr.Email).First(&emailValidation).RecordNotFound() {
		err := Model.ErrorMessage{"email", "Email must unique"}
		errors.Error = append(errors.Error, err)
	}
	phoneValidation := Model.Users{}
	if !db.Where("mobile_number = ?", usr.MobileNumber).First(&phoneValidation).RecordNotFound() {
		err := Model.ErrorMessage{"mobile_number", "Mobile number must unique"}
		errors.Error = append(errors.Error, err)
	}
	return errors
}
