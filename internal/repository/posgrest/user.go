package posgrest

import (
	"context"
	"fmt"

	"github.com/jeffleon1/ionix/pkg/entities"
	"gorm.io/gorm"
)

type userRepository[U entities.User] struct {
	repository[U]
}

func NewUserRepository[U entities.User](db *gorm.DB) entities.UserRepository[U] {
	return &userRepository[U]{repository: repository[U]{db}}
}

func (u *userRepository[U]) Login(ctx context.Context, email, password string) (*entities.User, error) {
	var user entities.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if !user.VerifyPassword(password) {
		return nil, fmt.Errorf("Invalid password")
	}

	return &user, nil
}

func (u *userRepository[U]) Signup(ctx context.Context, user *U) error {
	return u.Create(ctx, user)
}
