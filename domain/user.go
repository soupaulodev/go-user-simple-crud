package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetAll() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
