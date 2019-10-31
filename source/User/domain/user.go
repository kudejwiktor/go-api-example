package domain

type User struct {
	ID   int    `json:"id" DB:"id"`
	Name string `json:"name" DB:"name"`
}

type UserRepository interface {
	GetUserOfId(id int) (*User, error)
}
