package users

import (
	"errors"
	"fmt"
	"strings"

	"gihub.com/olamilekan000/bp-go/constants"
	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier()
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Type      string `json:"type"`
	Store     string `json:"store"`
}

func (u *User) Validate() error {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.New(constants.InavlidEmailAddress)
	}

	if err := u.verifyEmial(u.Email); err != nil {
		return errors.New(constants.InavlidEmailAddress)
	}

	return nil
}

func (u *User) verifyEmial(email string) error {
	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return nil
	}

	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return nil
	}

	return nil
}
