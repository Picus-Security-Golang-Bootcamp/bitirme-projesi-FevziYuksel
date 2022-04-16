package CartHandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

func (c *CartHandler) Dummy(context *gin.Context) {
	fmt.Println("unutma")
}

/*
import (
	UserHandler "FinalProjectGO/API/users"
	"FinalProjectGO/Models/cart"
	"FinalProjectGO/Models/product"
	"FinalProjectGO/pkg/config"
	jwt_helper "FinalProjectGO/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RequestBody struct {
	ID     uint `json:"id"`
	Amount int  `json:"amount"`
}

type CartHandler struct {
}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

var cfg *config.Config

func init() {
	cfg1, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg = cfg1

}

// AddProductToCart adds product to cart
/*
func (c *CartHandler) InsertProductToCart(productId uint, amount int, userId uint) (*product.Product, error) {

	//
	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}

	if chosenProduct.Stock <= amount {
		return nil, errors.New("ProductNotEnoughStockError")
	}

	if amount <= 0 {
		return nil, helpers.InvalidNumberOfProductsError
	}

	if IsProductExist(userId, productId) {
		return nil, errors.New("ProductAlreadyExistInCart")
	}

	return chosenProduct, nil
}

func (c *CartHandler) AddProductToCart(context *gin.Context) {
	//Body decode
	SecretKey := cfg.JWTConfig.SecretKey
	decodedToken := jwt_helper.VerifyToken(context.GetHeader("Authorization"), SecretKey)
	/*
		//err ekle middleware de aynı fonsiyon var
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			context.Abort()
			return
		}

	//Rewrite to body (memory address)
	var body RequestBody
	err := UserHandler.DecodeBody(&body, context)
	if err != nil {
		return
	}

	chosenProduct, err := c.CartService.AddProductToCart(body.ID, body.Amount, decodedToken.UserId)

	if err != nil {
		if err == helpers.ProductAlreadyExistInCart {
			cartDetails, _ := c.CartService.UpdateProductInCart(body.ID, decodedToken.UserId, body.Amount)
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

	// Update user main cart
	fmt.Println(decodedToken)
	cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

	// Add product to cart details
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
*/
