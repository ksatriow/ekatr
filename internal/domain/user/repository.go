package user

type Repository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByID(id int) (*User, error)
	FindAll() ([]*User, error)
	DeleteByID(id int) error
}
