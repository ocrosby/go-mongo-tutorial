package models

import (
	"fmt"
	"strings"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Slug() string {
	if areAllNameFieldsEmpty(u) {
		return ""
	}

	return strings.ToLower(fmt.Sprintf("%s-%s", u.FirstName, u.LastName))
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s %s", u.FirstName, u.LastName)
}

func areAllNameFieldsEmpty(u *User) bool {
	return u.FirstName == "" && u.LastName == ""
}
