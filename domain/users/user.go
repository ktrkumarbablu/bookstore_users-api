package users

type User struct {
	Id          int64
	FirstName   string `json:"id"`
	LastName    string `json:"firstName"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}
