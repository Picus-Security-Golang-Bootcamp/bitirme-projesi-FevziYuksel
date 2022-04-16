package ProductHandler

import (
	"FinalProjectGO/API/bodyDecoder"
	"FinalProjectGO/Models/category"
	"FinalProjectGO/Models/product"
	"FinalProjectGO/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

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
	//Check if category of the product exist if not don't add it
	if !category.IsCategoryExist(body.CategoryName) {
		context.JSON(http.StatusAlreadyReported, gin.H{
			//"message": helpers.CategoryNotFoundError.Error(),
			"message": "CategoryNotFoundError",
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
func (p *ProductHandler) ListProducts(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, allProducts := product.GetAllProducts(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, allProducts)

	if len(products) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			//"message": helpers.PageCouldNotBeFoundError.Error(),
			"message": "PageCouldNotBeFoundError",
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
func (p *ProductHandler) SearchProduct(context *gin.Context) {
	
}
