package API

import (
	CartHandler "FinalProjectGO/API/cart"
	CategoryHandler "FinalProjectGO/API/category"
	ProductHandler "FinalProjectGO/API/product"
	UserHandler "FinalProjectGO/API/users"
	Middleware "FinalProjectGO/pkg/middleware"
	"github.com/gin-gonic/gin"
)

/*
	//belki uygularım
	// Set envs for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
*/

func RegisterHandlers(r *gin.Engine) {

	userHandler := UserHandler.NewUserHandler()
	categoryHandler := CategoryHandler.NewCategoryHandler()
	cartHandler := CartHandler.NewCartHandler()
	productHandler := ProductHandler.NewProductHandler()

	userGroup := r.Group("/user")
	userGroup.POST("/signup", userHandler.CreateUser)
	userGroup.POST("/login", userHandler.Login)

	categoryGroup := r.Group("/category")
	categoryGroup.POST("/create", Middleware.AuthMiddleware(), categoryHandler.CreateBulkCategory)
	categoryGroup.GET("/list", categoryHandler.ListAllCategories)

	cartGroup := r.Group("/cart")
	cartGroup.POST("/add", Middleware.AuthMiddleware(), cartHandler.Dummy) //cartHandler.AddProductToCart

	/*
		cartGroup.GET("/list", Middleware.AuthForGeneral(), cartHandler.GetCartList)
		cartGroup.DELETE("/delete", Middleware.AuthForGeneral(), cartHandler.DeleteProductFromCart)
		cartGroup.PUT("/update", Middleware.AuthForGeneral(), cartHandler.UpdateProductInCart)

	*/
	productGroup := r.Group("/product")
	productGroup.POST("/create", Middleware.AuthMiddleware(), productHandler.CreateProduct)
	productGroup.GET("/list", productHandler.ListProducts)
	productGroup.GET("/search", productHandler.SearchProduct)
}
