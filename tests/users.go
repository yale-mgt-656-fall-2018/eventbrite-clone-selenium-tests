package tests

import (
	"math/rand"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

// User ...
//
type User struct {
	name     string
	email    string
	password string
	flaw     string
}

func randomUser() User {
	return User{
		name:     randomdata.FullName(randomdata.RandomGender),
		email:    randomString(7+rand.Intn(5)) + "@" + randomString(7+rand.Intn(5)) + ".com",
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
	u.flaw = "invalid email"
	users = append(users, u)

	u = randomUser()
	u.email = randomString(50) + "@gmail.com"
	u.flaw = "email that is more than 50 letters"
	users = append(users, u)

	u = randomUser()
	u.name = ""
	u.flaw = "empty name"
	users = append(users, u)

	u = randomUser()
	u.name = randomString(51)
	u.flaw = "name that is more than 50 letters"
	users = append(users, u)

	u = randomUser()
	u.password = ""
	u.flaw = "empty password"
	users = append(users, u)

	u = randomUser()
	u.password = randomString(51)
	u.flaw = "password that is more than 50 letters"
	users = append(users, u)

	return users
}
