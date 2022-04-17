package orderHandler

import (
	"FinalProjectGO/Models/cart"
	"FinalProjectGO/Models/order"
	"FinalProjectGO/Models/product"
	jwt_helper "FinalProjectGO/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (o *OrderHandler) GetOrderList(context *gin.Context) {
	decodedToken, _ := jwt_helper.VerifyToken(context.GetHeader("Authorization"))

	userOrders := order.FindUserOrders(decodedToken.UserId)
	if len(userOrders) == 0 {
		context.JSON(http.StatusOK, gin.H{
			"message": "OrderNotFoundError",
		})
		context.Abort()
		return
	}

	outPut := make([]Order, len(userOrders))
	for i, chosenOrder := range userOrders {
		outPut[i] = Order{
			TotalPrice: chosenOrder.TotalPrice,
			Amount:     chosenOrder.Amount,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Your orders": outPut,
	})
}

func (o *OrderHandler) GetOrderDetails(context *gin.Context) {
	decodedToken, _ := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	id, isOk := context.GetQuery("id")
	if !isOk {
		context.JSON(http.StatusBadRequest, errors.New("IdIsRequiredError"))
		context.Abort()
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, errors.New("InvalidIdError"))
		context.Abort()
		return
	}
	chosenOrder := order.SearchById(uint(orderId), decodedToken.UserId)

	if chosenOrder.ID == 0 {
		context.JSON(http.StatusBadRequest, errors.New("OrderNotFoundError"))
		context.Abort()
		return
	}
	orderDetails := order.FindOrderDetails(uint(orderId))
	outPut := make([]Product, len(orderDetails))
	for i, orderDetail := range orderDetails {
		outPut[i] = Product{
			ProductName: orderDetail.ProductName,
			Amount:      orderDetail.Amount,
			UnitPrice:   orderDetail.UnitPrice,
			TotalPrice:  orderDetail.TotalPrice,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Order Details": outPut,
	})

}

func (o *OrderHandler) CancelOrder(context *gin.Context) {
	decodedToken, _ := jwt_helper.VerifyToken(context.GetHeader("Authorization"))
	id, isOk := context.GetQuery("id")
	if !isOk {
		context.JSON(http.StatusBadRequest, errors.New("IdIsRequiredError"))
		context.Abort()
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, errors.New("InvalidIdError"))
		context.Abort()
		return
	}

	err = checkOrderTime(uint(orderId), decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		context.Abort()
		return
	}

	chosenOrderDetails := order.FindOrderDetails(uint(orderId))
	for _, chosenOrderDetail := range chosenOrderDetails {
		chosenProduct := product.SearchById(chosenOrderDetail.ProductId)
		chosenProduct.Stock += chosenOrderDetail.Amount
		product.Update(chosenProduct)
		order.DeleteModel(chosenOrderDetail)
	}

	order.DeleteOrder(uint(orderId))

	context.JSON(http.StatusOK, gin.H{
		"message": "Order has been canceled",
	})
}
func checkOrderTime(orderId uint, userId uint) error {
	checkOrder := order.SearchById(orderId, userId)
	if checkOrder.ID == 0 {
		return errors.New("OrderNotFoundError")
	}
	currentTime := time.Now()
	timeDifference := currentTime.Sub(currentTime).Hours()

	if timeDifference > 24*14 {
		return errors.New("OrderCancelError")
	}
	return nil
}

func (o *OrderHandler) CreateOrder(context *gin.Context) {

	decodedToken, _ := jwt_helper.VerifyToken(context.GetHeader("Authorization"))

	allProductsInCart := cart.GetAllCartDetailsOfUser(decodedToken.UserId)
	if len(*allProductsInCart) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("CartIsEmptyError"),
		})
		context.Abort()
		return
	}

	var outOfStockProducts []string
	for _, productInCart := range *allProductsInCart {
		chosenProduct := product.SearchById(productInCart.ProductId)
		if chosenProduct.Stock < productInCart.Amount {
			newMessage := fmt.Sprintf("Product %s has no enought stock. Available stock is %d, your amount is %d",
				chosenProduct.ProductName, chosenProduct.Stock, productInCart.Amount)
			outOfStockProducts = append(outOfStockProducts, newMessage)
		}
	}
	if len(outOfStockProducts) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":  "NotEnoughStockError",
			"products": outOfStockProducts,
		})
		context.Abort()
		return
	}

	userCart := cart.SearchById(decodedToken.UserId)
	newOrder := order.NewOrder(userCart.TotalPrice, userCart.Amount, userCart.UserId)
	order.CreateOrderTable(newOrder)
	userCart.TotalPrice = 0
	userCart.Amount = 0
	cart.Update(userCart)

	for _, productInCart := range *allProductsInCart {
		chosenProduct := product.SearchById(productInCart.ProductId)
		newStock := chosenProduct.Stock - productInCart.Amount
		order.CreateOrderDetailTable(newOrder.ID, chosenProduct.ID, productInCart.Amount, chosenProduct.Price, productInCart.TotalPrice, chosenProduct.ProductName)
		product.UpdateStock(*chosenProduct, newStock)
		cart.DeleteModel(&productInCart)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
	})
}
