package tokenutil

import (
	"fmt"
	"team-work-space/domain"

	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetDataFromClaims(requestToken string, secret string) (*domain.JwtData, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("Invalid token:")
	}
	jwtData := &domain.JwtData{
		UserId:   claims["userId"].(string),
		Username: claims["username"].(string),
		RoleId:   claims["roleId"].(string),
	}
	return jwtData, nil
}
