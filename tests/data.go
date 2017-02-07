package tests

import (
	"math/rand"

	"github.com/Pallinder/go-randomdata"
	"github.com/yale-cpsc-213/social-todo-selenium-tests/tests/selectors"
)

// User ...
//
type User struct {
	name        string
	email       string
	password    string
	description string
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
func (u User) registerFormData() map[string]string {
	data := map[string]string{
		selectors.RegisterFormEmail:                u.email,
		selectors.RegisterFormName:                 u.name,
		selectors.RegisterFormPassword:             u.password,
		selectors.RegisterFormPasswordConfirmation: u.password,
	}
	return data
}

func getBadUsers() []User {
	var u User
	var users []User

	u = randomUser()
	u.email = "sdfsdfsd"
	u.description = "invalid email"
	users = append(users, u)

	u = randomUser()
	u.email = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx@gmail.com"
	u.description = "email that is more than 50 letters"
	users = append(users, u)

	u = randomUser()
	u.name = ""
	u.description = "empty name"
	users = append(users, u)

	u = randomUser()
	u.name = randomString(51)
	u.description = "name that is more than 50 letters"
	users = append(users, u)

	u = randomUser()
	u.password = ""
	u.description = "empty password"
	users = append(users, u)

	u = randomUser()
	u.password = randomString(51)
	u.description = "password that is more than 50 letters"
	users = append(users, u)

	return users
}
