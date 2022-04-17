package API

import (
	CartHandler "FinalProjectGO/API/cart"
	CategoryHandler "FinalProjectGO/API/category"
	OrderHandler "FinalProjectGO/API/order"
	ProductHandler "FinalProjectGO/API/product"
	UserHandler "FinalProjectGO/API/users"
	Middleware "FinalProjectGO/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {

	userHandler := UserHandler.NewUserHandler()
	userGroup := r.Group("/user")
	userGroup.POST("/signup", userHandler.CreateUser)
	userGroup.POST("/login", userHandler.Login)

	categoryHandler := CategoryHandler.NewCategoryHandler()
	categoryGroup := r.Group("/category")
	categoryGroup.POST("/create", Middleware.AdminCheck(), categoryHandler.CreateBulkCategory)
	categoryGroup.GET("/list", categoryHandler.ListAllCategories)

	productHandler := ProductHandler.NewProductHandler()
	productGroup := r.Group("/product")
	productGroup.POST("/create", Middleware.AdminCheck(), productHandler.CreateProduct)
	productGroup.GET("/list", productHandler.ListProducts)
	productGroup.GET("/search", productHandler.SearchProduct)
	//Test Hepsini
	productGroup.DELETE("/delete", Middleware.AdminCheck(), productHandler.DeleteProduct) //çalışıyor gibi
	productGroup.PUT("/update", Middleware.AdminCheck(), productHandler.UpdateProduct)    //sorunlu

	cartHandler := CartHandler.NewCartHandler()
	cartGroup := r.Group("/cart")
	cartGroup.POST("/add", Middleware.GeneralCheck(), cartHandler.AddProductToCart)
	cartGroup.GET("/list", Middleware.GeneralCheck(), cartHandler.GetCartList)
	cartGroup.DELETE("/delete", Middleware.GeneralCheck(), cartHandler.DeleteProductFromCart)
	cartGroup.PUT("/update", Middleware.GeneralCheck(), cartHandler.UpdateProductInCart)

	orderHandler := OrderHandler.NewOrderHandler()
	orderGroup := r.Group("/order")
	orderGroup.GET("/list", Middleware.GeneralCheck(), orderHandler.GetOrderList)
	orderGroup.POST("/create", Middleware.GeneralCheck(), orderHandler.CreateOrder)
	orderGroup.GET("/detail", Middleware.GeneralCheck(), orderHandler.GetOrderDetails)
	orderGroup.DELETE("/cancel", Middleware.GeneralCheck(), orderHandler.CancelOrder)

}
