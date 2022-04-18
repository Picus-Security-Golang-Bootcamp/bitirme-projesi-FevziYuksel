package UserHandler

import (
	"FinalProjectGO/API/bodyDecoder"
	"FinalProjectGO/Models/cart"
	User "FinalProjectGO/Models/users"
	jwt_helper "FinalProjectGO/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type Users struct {
	Email    string
	Password string
}

// CreateUser godoc
// @Summary Sign in to system with username and password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param RequestBody body Users true "Sign in info"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/signup [post]
func (h *UserHandler) CreateUser(context *gin.Context) {
	var body User.Users
	err := bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		return
	}

	newUser := User.NewUser(body.GetEmail(), body.GetPassword(), body.GetRole())

	if err := h.CheckUser(newUser); err != nil {
		context.JSON(http.StatusSeeOther, gin.H{
			"message": err.Error(),
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
	//New Cart for each User
	cart.CreateCardTable(cart.NewCart(newUser.ID))
}
func (h *UserHandler) CheckUser(user *User.Users) error {
	if User.IsUserExist(user.GetEmail()) {
		return errors.New("UserExistsError")
	}
	return nil
}

// Login godoc
// @Summary login to system with username and password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param RequestBody body Users true "Login info"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/login [post] //router
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
		return nil, errors.New("user is not found")
	}

	if users[0].GetPassword() != body.GetPassword() {
		return nil, errors.New("invalid password")
	}
	return &users[0], nil
}
