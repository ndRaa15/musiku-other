package jwt

import (
	"os"
	"time"

	"github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/golang-jwt/jwt/v4"
)

func EncodeToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
	})
	signedToken, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", errors.ErrSigningJWT
	}
	return signedToken, nil
}

func DecodeToken(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)
	if !ok || !decoded.Valid {
		return nil, errors.ErrClaimsJWT
	}
	return claims, nil
}
