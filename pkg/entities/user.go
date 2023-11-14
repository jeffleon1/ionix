package entities

import (
	"context"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type UserRepository[T User] interface {
	Repository[T]
	Login(ctx context.Context, username, password string) (*User, error)
	Signup(ctx context.Context, user *T) error
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
