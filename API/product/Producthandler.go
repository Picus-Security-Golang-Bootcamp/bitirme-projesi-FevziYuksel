package ProductHandler

import (
	"FinalProjectGO/API/bodyDecoder"
	"FinalProjectGO/Models/cart"
	"FinalProjectGO/Models/category"
	"FinalProjectGO/Models/product"
	"FinalProjectGO/pkg/pagination"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

type Product struct {
	ProductName  string
	Price        float64
	Stock        int
	CategoryName string
	SKU          string
}

// CreateProduct godoc
// @Summary Create a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param RequestBody body Product false "Product"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authentication"
// @Router /product/create [post]
func (p *ProductHandler) CreateProduct(context *gin.Context) {
	var body product.Product

	err := bodyDecoder.DecodeBody(&body, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	if product.IsProductExist(body.ProductName, body.SKU) {
		context.JSON(http.StatusAlreadyReported, gin.H{
			"message": "Product is already exist",
		})
		context.Abort()
		return
	}

	if !category.IsCategoryExist(body.CategoryName) {
		context.JSON(http.StatusAlreadyReported, gin.H{
			"message": "category not found",
		})
		context.Abort()
		return
	}
	categoryId := category.GetCategoryId(body.CategoryName)

	body.CategoryId = categoryId

	newProduct := product.NewProduct(body)

	product.CreateProduct(newProduct)

	context.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})
}

// ListProducts godoc
// @Summary List all products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /product/list [get]
func (p *ProductHandler) ListProducts(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, allProducts := product.GetAllProducts(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, allProducts)

	if len(products) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Page could not be found",
		})
		context.Abort()
		return
	}

	output := make([]ProductPage, len(products))
	for i, eachProduct := range products {
		output[i] = ProductPage{
			ProductId:    eachProduct.ID,
			ProductName:  eachProduct.ProductName,
			Price:        eachProduct.Price,
			Stock:        eachProduct.Stock,
			CategoryName: eachProduct.CategoryName,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Info":     paginatedResult,
		"Products": output,
	})
}

// SearchProduct godoc
// @Summary Search a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param search query string false "search"
// @Param page query int false "Page Index"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /product/search [get]
func (p *ProductHandler) SearchProduct(context *gin.Context) {

	search, isOk := context.GetQuery("search")
	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": "search parameter is required",
		})
		context.Abort()
		return
	}
	allProducts := product.SearchProduct(search)

	if len(allProducts) == 0 {
		context.JSON(http.StatusOK, gin.H{
			"message": "product not found",
		})
		context.Abort()
		return
	}

	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products := product.SearchProductWithPagination(search, pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, len(allProducts))

	if len(products) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Page could not be found",
		})
		context.Abort()
		return
	}

	output := make([]ProductPage, len(products))
	for i, eachProduct := range products {
		output[i] = ProductPage{
			ProductId:    eachProduct.ID,
			ProductName:  eachProduct.ProductName,
			Price:        eachProduct.Price,
			Stock:        eachProduct.Stock,
			CategoryName: eachProduct.CategoryName,
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"Info":     paginatedResult,
		"products": output,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /product/delete [delete]
func (p *ProductHandler) DeleteProduct(context *gin.Context) {
	productId, isOk := context.GetQuery("id")

	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": "InvalidIdError",
		})
		context.Abort()
		return
	}

	id, err := strconv.Atoi(productId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "InvalidIdError",
		})
		context.Abort()
		return
	}

	chosenProduct := product.SearchById(uint(id))
	if chosenProduct.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "ProductNotFoundError",
		})
		context.Abort()
		return
	}

	cartsHasChosenProduct := cart.GetCartDetailsByProductId(chosenProduct.ID)
	cartIsEmpty := len(*cartsHasChosenProduct) == 0
	if !cartIsEmpty {
		for _, cartDetail := range *cartsHasChosenProduct {
			newTotalPrice := -cartDetail.TotalPrice
			newAmount := -cartDetail.Amount
			cart.UpdateUserCart(cartDetail.CartId, newAmount, newTotalPrice)
			cart.DeleteProductInCart(cartDetail.CartId, chosenProduct.ID)
		}
	}

	product.DeleteProduct(*chosenProduct)

	context.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// UpdateProduct godoc
// @Summary Update a product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id query int false "id"
// @Param newStock query int true "stock"
// @Param newPrice query int true "price"
// @Param newName query string true "name"
// @Param newSku query string true "sku"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /product/update [put]
func (p *ProductHandler) UpdateProduct(context *gin.Context) {
	id, isProductId := context.GetQuery("id")
	newStock, isStock := context.GetQuery("stock")
	newPrice, isPrice := context.GetQuery("price")
	newName, isName := context.GetQuery("name")
	newSku, isSku := context.GetQuery("sku")

	if !isProductId {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "IdIsRequiredError",
		})
		context.Abort()
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "InvalidIdError",
		})
		context.Abort()
		return
	}

	chosenProduct := product.SearchById(uint(productId))
	if chosenProduct.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "ProductNotFoundError",
		})
		context.Abort()
		return
	}
	if isStock {
		stock, err := strconv.Atoi(newStock)
		if err != nil || stock <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "InvalidStockError",
			})
			context.Abort()
			return
		}
		product.UpdateStock(*chosenProduct, stock)

	}
	if isSku {
		if chosenProduct.SKU == newSku {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": "SkuIsSameError",
			})
			context.Abort()
			return
		}

		sameSkuProduct := product.SearchBySKU(newSku)
		if sameSkuProduct.ID != 0 {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": "SkuAlreadyExistError",
			})
			context.Abort()
			return
		}
		product.UpdateSKU(*chosenProduct, newSku)
	}

	cartsHasChosenProduct := cart.GetCartDetailsByProductId(chosenProduct.ID)
	cartIsEmpty := len(*cartsHasChosenProduct) == 0

	if isPrice {
		price, err := strconv.ParseFloat(newPrice, 64)
		if err != nil || price <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": errors.New("InvalidPriceError"),
			})
			context.Abort()
			return
		}

		if cartIsEmpty {
			for _, cartDetail := range *cartsHasChosenProduct {
				newTotalPrice := (chosenProduct.Price - price) * float64(cartDetail.Amount)
				newAmount := cartDetail.Amount

				cartDetail.UnitPrice = price
				cartDetail.TotalPrice = newTotalPrice
				cart.UpdateUserCart(cartDetail.CartId, newAmount, newTotalPrice)
				cart.UpdateModel(&cartDetail)
			}
		}

		product.UpdatePrice(*chosenProduct, price)
	}

	if isName {
		if chosenProduct.ProductName == newName {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": errors.New("NameIsSameError"),
			})
			context.Abort()
			return
		}

		sameNameProduct := product.SearchByProductName(newName)
		if sameNameProduct.ID != 0 {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": errors.New("NameAlreadyExistError"),
			})
			context.Abort()
			return
		}

		if cartIsEmpty {
			for _, cartDetail := range *cartsHasChosenProduct {
				cartDetail.ProductName = newName
				cart.UpdateModel(&cartDetail)
			}
		}
		product.UpdateName(*chosenProduct, newName)
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "product updated",
	})

}
