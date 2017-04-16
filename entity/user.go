package entity

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
