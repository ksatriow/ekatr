package user

import "time"

type UserID int

type UserType string

const (
	Owner  UserType = "owner"
	Kasir  UserType = "kasir"
	Pembeli UserType = "pembeli"
)

type User struct {
	ID           UserID
	Username     string
	Password     string
	Email        string
	Type         UserType
	ProfilePhoto string
	CreatedAt    time.Time
}

func NewUser(username, password, email string, userType UserType, profilePhoto string) *User {
	return &User{
		Username:     username,
		Password:     password,
		Email:        email,
		Type:         userType,
		ProfilePhoto: profilePhoto,
		CreatedAt:    time.Now(),
	}
}
