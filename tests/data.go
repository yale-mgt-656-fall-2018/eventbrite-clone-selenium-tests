package tests

import (
	"math/rand"

	"github.com/Pallinder/go-randomdata"
	"github.com/yale-cpsc-213/social-todo-selenium-tests/tests/selectors"
)

// User ...
//
type User struct {
	name     string
	email    string
	password string
}

func randomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func randomUser() User {
	return User{
		name:     randomdata.FullName(randomdata.RandomGender),
		email:    randomdata.Email(),
		password: randomString(10),
	}
}

func (u User) loginFormData() map[string]string {
	data := map[string]string{
		selectors.LoginFormEmail:    u.email,
		selectors.LoginFormPassword: u.password,
	}
	return data
}
