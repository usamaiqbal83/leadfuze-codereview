package users

import "github.com/pkg/errors"

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
}

// return combination in format like usamaiqbal
func (user *User) FullName() string {
	return user.FirstName + user.LastName
}

// returns initials from first name and last name
func (user *User) Initials() string {
	a := user.firstLetter(user.FirstName)
	b := user.firstLetter(user.LastName)
	return a + b
}

// return combination in format like uiqbal
func (user *User) Combination1() string {
	a := user.firstLetter(user.FirstName)
	b := user.LastName
	return a + b
}

// return combination in format like usamai
func (user *User) Combination2() string {
	a := user.FirstName
	b := user.firstLetter(user.LastName)
	return a + b
}

// return combination in format like
func (user *User) Combination3(infix string) (string, error) {
	if infix == "" {
		return "", errors.New("infix can not be empty")
	}

	if user.FirstName == "" {
		return user.LastName, nil
	}

	if user.LastName == "" {
		return user.FirstName, nil
	}

	return user.FirstName + infix + user.LastName, nil
}

// private methods
func (user *User) firstLetter(src string) string {
	if src == "" {
		return src
	}

	return src[0:1]
}