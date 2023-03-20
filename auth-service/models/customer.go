package models

type Customer struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
	Mobile string `json:"mobile" db:"mobile"`
}
