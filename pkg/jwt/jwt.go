package jwt_helper

import (
	"FinalProjectGO/Models/users"
	config "FinalProjectGO/pkg/config"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type DecodedToken struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Iat    int    `json:"iat"`
}

var (
	cfg       *config.Config
	SecretKey string
)

func init() {
	cfg1, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg = cfg1
	SecretKey = cfg.JWTConfig.SecretKey
}

// CreateToken creates a new token
func GenerateToken(user *users.Users) string {

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"role":   user.Role,
		"iat":    time.Now().Unix(),
		"exp": time.Now().Add(100 *
			time.Hour).Unix(),
	})

	hmacSecretString := SecretKey
	hmacSecret := []byte(hmacSecretString)
	token, _ := jwtClaims.SignedString(hmacSecret)
	return token
}
func VerifyToken(token string) (*DecodedToken, error) {
	hmacSecretString := SecretKey
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, errors.New("DecoderError")
	}
	if decoded.Valid {
		decodedClaims := decoded.Claims.(jwt.MapClaims)

		var decodedToken DecodedToken
		jsonString, _ := json.Marshal(decodedClaims)
		err = json.Unmarshal(jsonString, &decodedToken)
		if err != nil {
			return nil, err
		}
		return &decodedToken, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("ExpiredTokenError")
		}
	} else {
		return nil, errors.New("InvalidTokenError")

	}
	return nil, err
}
