package mw

import (
	jwtHelper "FinalProjectGO/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeneralCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("Authorization") != "" {
			_, err := jwtHelper.VerifyToken(context.GetHeader("Authorization"))
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
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
