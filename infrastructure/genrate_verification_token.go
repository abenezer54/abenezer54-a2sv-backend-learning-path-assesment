package infrastrucutre

import (
	"loan-api/domain"
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateVerificationToken(ctx context.Context, user domain.User, tokenSecret string, tokenExpiry int) (string, error) {
	claims := jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"password":  user.Password,
		"exp":       time.Now().Add(time.Hour * time.Duration(tokenExpiry)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
