package bodyDecoder

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var val = validator.New()

type ApiError struct {
	Field string
	Msg   string
}

func DecodeBody(body interface{}, context *gin.Context) error {
	if context.Request.Header.Get("Content-Type") != "application/json" {
		return errors.New("UnsupportedMediaType")
		//return UnsupportedMediaType
	}

	if err := context.ShouldBindJSON(body); err != nil {
		return err
	}

	if err := val.Struct(body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
			context.Abort()
			return errors.New("InvalidFieldError")
			//return InvalidFieldError
		}
		return err
	}

	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "This field must be at least 8 characters"
	case "max":
		return "This field must be at most 20 characters"
	case "number":
		return "This field must be a number"

	case "gt":
		return "This field must be greater than 1"
	}

	return ""
}
