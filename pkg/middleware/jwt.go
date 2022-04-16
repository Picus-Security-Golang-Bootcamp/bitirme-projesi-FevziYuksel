package mw

import (
	"FinalProjectGO/pkg/config"
	jwtHelper "FinalProjectGO/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
)

var cfg *config.Config

func AuthMiddleware() gin.HandlerFunc {
	cfg1, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg = cfg1
	SecretKey := cfg.JWTConfig.SecretKey
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), SecretKey)
			
			if decodedClaims != nil {
				if decodedClaims.Role == "admin" {
					c.Next()
					c.Abort()
					return
					return
				}
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
