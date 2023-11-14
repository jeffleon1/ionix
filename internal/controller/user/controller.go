package user

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeffleon1/ionix/pkg/entities"
	"github.com/jeffleon1/ionix/pkg/models"
)

type Controller interface {
	Login(context.Context, string, string) (string, error)
	Signup(context.Context, entities.User) (string, error)
}

type controller struct {
	repo          entities.UserRepository[entities.User]
	secretKey     string
	tokenDuration int
}

func New(repo entities.UserRepository[entities.User], secretKey string, tokenDuration int) Controller {
	return &controller{
		repo,
		secretKey,
		tokenDuration,
	}
}

func (c *controller) Login(ctx context.Context, email, password string) (string, error) {
	user, err := c.repo.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	return c.generateJWT(int(user.ID))

}

func (c *controller) Signup(ctx context.Context, user entities.User) (string, error) {
	if err := user.HashPassword(); err != nil {
		return "", err
	}

	if err := c.repo.Signup(ctx, &user); err != nil {
		return "", err
	}

	return c.generateJWT(int(user.ID))
}

func (c *controller) generateJWT(userID int) (string, error) {
	claims := models.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
