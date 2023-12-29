package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserUpdateInput struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}
