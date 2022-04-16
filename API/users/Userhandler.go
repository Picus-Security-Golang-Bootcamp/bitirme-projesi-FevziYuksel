package UserHandler

import (
	"FinalProjectGO/API/bodyDecoder"
	User "FinalProjectGO/Models/users"
	jwt_helper "FinalProjectGO/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

//Constructor
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

//Signup block

//CheckUser Bloğunu kontrol et sonra sil
func (h *UserHandler) CheckUser(user *User.Users) error {
	if User.IsUserExist(user.GetEmail()) {
		//return helpers.UserExistsError
		return errors.New("helpers.UserExistsError")
	}
	return nil
	//newUser'ın hemen altına
	//CheckUser Bloğunu kontrol et sonra sil
	/*
		if err := h.CheckUser(newUser); err != nil {
			context.JSON(http.StatusSeeOther, gin.H{
				"message": err.Error(),
			})
			context.Abort()
			return
		}

	*/
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(context *gin.Context) {
	var body User.Users
	err := bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		return
	}

	newUser := User.NewUser(body.GetEmail(), body.GetPassword(), body.GetRole())

	if User.IsUserExist(newUser.GetEmail()) {
		//return helpers.UserExistsError
		context.JSON(http.StatusSeeOther, gin.H{
			"message": errors.New("helpers.UserExistsError"),
		})
		context.Abort()
		return
	}

	err = User.CreateUser(newUser)
	if err != nil {
		return
	}

	token := jwt_helper.GenerateToken(newUser)
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
	})

	// Creates cart for new user
	//cart.Create(cart.NewCartRepository(newUser.ID))
}

//Login block

func (h *UserHandler) Login(context *gin.Context) {
	var body User.Users

	err := bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	loggedUser, err := h.CheckLogin(body)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		context.Abort()
		return
	}
	token := jwt_helper.GenerateToken(loggedUser)
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}
func (h *UserHandler) CheckLogin(body User.Users) (*User.Users, error) {
	users := User.SearchByEmail(body)
	fmt.Println(users)
	if len(users) == 0 {
		//return nil, helpers.UserNotFoundError
		return nil, errors.New("UserNotFoundError")
	}

	if users[0].GetPassword() != body.GetPassword() {
		//return nil, helpers.InvalidPasswordError
		return nil, errors.New("InvalidPasswordError")
	}
	return &users[0], nil
}
