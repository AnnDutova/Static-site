package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example5.org",
		Password: "password",
	}
}
