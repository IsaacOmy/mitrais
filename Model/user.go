package Model

type Users struct {
	Id           int `gorm:"primary_key"`
	Email        string
	Password     string
	FirstName    string
	LastName     string
	DateOfBirth  string
	Gender       string
	MobileNumber string
}
