package jwttoken

import (
	"encoding/base64"
	"fmt"
	"go-project/config"

	"github.com/golang-jwt/jwt"
)

func ValidateTokenHeader(token string) (string, error) {
	decoderPublicKey, err := base64.StdEncoding.DecodeString(config.AppEnv.AccTokenPublicKey)
	if err != nil {
		return " ", fmt.Errorf("could not decode: %w", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decoderPublicKey)
	if err != nil {
		return " ", fmt.Errorf("validate parse key: %w", err)
	}
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("ini oke", ok)

			return nil, fmt.Errorf("cnexpacted Method: %s", t.Header["exp"])
		}
		return key, nil
	})
	if err != nil {

		return " ", fmt.Errorf("validate :%w", err)
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return " ", fmt.Errorf("validate:invalid token")
	}
	return claims["user_id"].(string), nil
}
