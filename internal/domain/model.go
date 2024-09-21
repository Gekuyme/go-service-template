package domain

type User struct {
	Name     string `json:"name"`
	Email    int    `json:"email"`
	Password string `json:"password"`
}
