package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedding example
type Admin struct {
	email    string
	password string
	User     // it can be anonymous field. when it is anonymous, you access its attributes and methods directly like if they were part of Admin
}

// ACCESOR method. It works with values or pointers, user or *user
func (u User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// MUTATOR method. It works with pointers, *user
func (u *User) Clear() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{ // if was anonymous field: User else: user: user
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "01/01/1970",
			createdAt: time.Now(),
		},
	}
}

// CONSTRUCTOR method. It works with values or pointers, user and return user OR *user and return &user
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
