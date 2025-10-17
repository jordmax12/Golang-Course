package user

import (
	"fmt"
	commands "main/commands"
	validations "main/validations"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
	Age       int
}

type Admin struct {
	email string
	password string
	User
}

func ChangeUserToAdmin(user *User) (*Admin, error) {
	email, err := commands.PromptUntilValid("Please enter your email: ", validations.ValidateEmail)
	if err != nil {
		return nil, err
	}
	
	password, err := commands.PromptUntilValid("Please enter your password: ", validations.ValidatePassword)
	if err != nil {
		return nil, err
	}
	
	return &Admin{
		email: email,
		password: password,
		User: *user,
	}, nil
}

func (a *Admin) String() string {
	return fmt.Sprintf("Email: %s\nPassword: %s\nUser: %s",
		a.email, a.password, a.User.String())
}

func New() (*User, error) {
	firstName, err := commands.PromptUntilValid("Please enter your first name: ", validations.ValidateNonEmpty)
	if err != nil {
		return nil, err
	}
	
	lastName, err := commands.PromptUntilValid("Please enter your last name: ", validations.ValidateNonEmpty)
	if err != nil {
		return nil, err
	}
	
	birthdate, err := commands.PromptUntilValid("Please enter your birthdate (MM/DD/YYYY): ", validations.ValidateBirthdate)
	if err != nil {
		return nil, err
	}
	
	parsedDate, _ := time.Parse("01/02/2006", birthdate)
	
	u := &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now().UTC(),
		Age:       time.Now().Year() - parsedDate.Year(),
	}
	
	return u, nil
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s %s\nBirthdate: %s\nAge: %d\nCreated: %s",
		u.FirstName, u.LastName, u.Birthdate, u.Age, u.CreatedAt.Format(time.RFC3339))
}
