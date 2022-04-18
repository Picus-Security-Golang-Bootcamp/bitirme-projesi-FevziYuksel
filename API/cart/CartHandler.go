package CartHandler

import (
	"FinalProjectGO/API/bodyDecoder"
	"FinalProjectGO/Models/cart"
	jwt_helper "FinalProjectGO/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

type Carts struct {
}

// AddProductToCart godoc
// @Summary Add a product to cart
// @Tags Carts
// @Accept  json
// @Produce  json
// @Param RequestBody body RequestBody true "Product"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /cart/add [post]
func (c *CartHandler) AddProductToCart(context *gin.Context) {

	var body RequestBody
	decodedToken, err := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	err = bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		return
	}
	fmt.Println(body)
	chosenProduct, errCheck, err := AddProductToCart(body.ID, body.Amount, decodedToken.UserId)

	if err != nil {

		if errCheck {

			cartDetails, _ := UpdateProductInCart(body.ID, decodedToken.UserId, body.Amount)
			newAmount := body.Amount + cartDetails.Amount
			newTotalPrice := cartDetails.TotalPrice + (float64(newAmount) * cartDetails.UnitPrice)

			cart.UpdateProductInCart(decodedToken.UserId, cartDetails.ProductId, newAmount, newTotalPrice)
			cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

			context.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
			context.Abort()
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	newTotalPrice := chosenProduct.Price * float64(body.Amount)

	cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

	cart.CreateCartDetails(&cart.CartDetails{
		ProductName: chosenProduct.ProductName,
		CartId:      decodedToken.UserId,
		ProductId:   chosenProduct.ID,
		Amount:      body.Amount,
		UnitPrice:   chosenProduct.Price,
		TotalPrice:  newTotalPrice,
	})

	context.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})

}

// GetCartList godoc
// @Summary List card contents
// @Tags Carts
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /cart/list [get]
func (c *CartHandler) GetCartList(context *gin.Context) {
	decodedToken, err := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	allCartDetails, err := GetCartList(decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	outPut := make([]Product, len(*allCartDetails))
	var TotalPrice float64
	for i, cartDetail := range *allCartDetails {
		TotalPrice += cartDetail.TotalPrice
		outPut[i] = Product{
			ProductName: cartDetail.ProductName,
			Amount:      cartDetail.Amount,
			UnitPrice:   cartDetail.UnitPrice,
			TotalPrice:  cartDetail.TotalPrice,
			ProductId:   cartDetail.ProductId,
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "Total Price": TotalPrice, "data": outPut})

}

// DeleteProductFromCart godoc
// @Summary Delete a product in cart
// @Tags Carts
// @Accept  json
// @Produce  json
// @Param id query int false "id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /cart/delete [delete]
func (c *CartHandler) DeleteProductFromCart(context *gin.Context) {
	decodedToken, err := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	id, isOk := context.GetQuery("id")

	if !isOk {
		context.JSON(http.StatusBadRequest, gin.H{"message": "IdIsRequiredError"})
		context.Abort()
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "InvalidIdError"})
		context.Abort()
		return
	}

	cartDetails, err := DeleteProductFromCart(decodedToken.UserId, uint(productId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	newAmount := -cartDetails.Amount
	newTotalPrice := -cartDetails.TotalPrice

	cart.UpdateUserCart(decodedToken.UserId, newAmount, newTotalPrice)

	cart.DeleteProductInCart(decodedToken.UserId, uint(productId))

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted from cart"})
}

// UpdateProductInCart godoc
// @Summary Update a product in cart
// @Tags Carts
// @Accept  json
// @Produce  json
// @Param RequestBody body RequestBody true "Product"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /cart/update [put]
func (c *CartHandler) UpdateProductInCart(context *gin.Context) {
	var body RequestBody
	decodedToken, err := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	err = bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	cartDetails, err := UpdateProductInCart(body.ID, decodedToken.UserId, body.Amount)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	var newTotalPrice float64
	if cartDetails.Amount > body.Amount {
		difference := cartDetails.Amount - body.Amount
		newTotalPrice = cartDetails.TotalPrice - (float64(difference) * cartDetails.UnitPrice)
	} else {
		newTotalPrice = float64(body.Amount) * cartDetails.UnitPrice
	}

	cart.UpdateProductInCart(decodedToken.UserId, cartDetails.ProductId, body.Amount, newTotalPrice)

	cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

	context.JSON(http.StatusOK, gin.H{"message": "Product updated in cart"})

}
