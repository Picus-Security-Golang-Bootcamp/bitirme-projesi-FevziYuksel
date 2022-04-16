package jwt_helper

import (
	Users "FinalProjectGO/Models/users"
	config "FinalProjectGO/pkg/config"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

//struct'ı düzelt rolü modifiye et ??

type DecodedToken struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Iat    int    `json:"iat"`
}

/*
func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}
*/

// CreateToken creates a new token
func GenerateToken(user *Users.Users) string {

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"role":   user.Role,
		"iat":    time.Now().Unix(),
		"exp": time.Now().Add(100 *
			time.Hour).Unix(),
	})
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	SecretKey := cfg.JWTConfig.SecretKey

	hmacSecretString := SecretKey
	hmacSecret := []byte(hmacSecretString)
	token, _ := jwtClaims.SignedString(hmacSecret)
	return token
}

//Error handling kısmı yok eklemelimiyim ?
func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	_ = json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken
}
