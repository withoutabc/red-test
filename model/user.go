package model

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Management struct {
	Repository string `json:"repository"`
	Name       string `json:"name"`
}
