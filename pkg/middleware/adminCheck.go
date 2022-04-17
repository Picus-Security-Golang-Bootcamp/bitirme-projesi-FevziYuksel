package mw

import (
	jwtHelper "FinalProjectGO/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("Authorization") != "" {
			decodedClaims, err := jwtHelper.VerifyToken(context.GetHeader("Authorization"))
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				context.Abort()
				return
			}
			if decodedClaims.Role != "admin" {
				context.JSON(http.StatusForbidden, gin.H{
					"error": errors.New("you are not allowed to use this endpoint"),
				})
				context.Abort()
				return
			} else {
				context.Next()
				context.Abort()
				return
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": errors.New("you are not authorized"),
			})
			context.Abort()
			return
		}
	}
}
