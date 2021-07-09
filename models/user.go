package models

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}
