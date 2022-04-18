package CategoryHandler

import (
	"FinalProjectGO/CSV"
	Category "FinalProjectGO/Models/category"
	"FinalProjectGO/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

type category struct {
	Name string
}

// CreateBulkCategory godoc
// @Summary Create and add categories from csv file
// @Tags category
// @Accept  json
// @Produce  json
// @Param   	 file formData file true "Category CSV"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /category/create [post]
func (c *CategoryHandler) CreateBulkCategory(context *gin.Context) {

	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(400, gin.H{
			"message": "FileError",
		})
		context.Abort()
		return
	}
	err = context.SaveUploadedFile(file, "CSV/"+file.Filename)

	if err != nil {
		context.JSON(400, gin.H{
			"Error1": err.Error(),
		})
		context.Abort()
		return
	}

	categoryList, err := CSV.CSVtoCategory("CSV/" + file.Filename)

	if err != nil {
		context.JSON(400, gin.H{
			"Error2": err.Error(),
		})
		context.Abort()
		return
	}

	var (
		newCategories []Category.Category
		existedNames  []string
	)
	for _, model := range categoryList {
		if !Category.IsCategoryExist(model.Name) {
			newCategories = append(newCategories, model)
		} else {
			existedNames = append(existedNames, model.GetName())
		}
	}
	if len(existedNames) > 0 {
		context.JSON(http.StatusAlreadyReported, gin.H{
			"message":            "Category already exist others are created",
			"existCategoryNames": existedNames,
		})
		context.Abort()
		return
	}
	if len(newCategories) != 0 {
		_ = Category.CreateCategoryTable(newCategories)
		_ = CSV.CSVtoProduct("CSV/Product.csv")
		context.JSON(http.StatusOK, gin.H{
			"message": "Category created successfully",
		})
	}
}

// ListAllCategories godoc
// @Summary List all product categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /category/list [get]
func (c *CategoryHandler) ListAllCategories(context *gin.Context) {

	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	categories, allCategories := Category.GetAllCategories(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, allCategories)

	if len(categories) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Page could not be found",
		})
		context.Abort()
		return
	}

	output := make([]CategoryPage, len(categories))
	for i, eachCategory := range categories {
		output[i] = CategoryPage{
			CategoryId:   eachCategory.ID,
			CategoryName: eachCategory.GetName(),
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Info":     paginatedResult,
		"Products": output,
	})

}

func (c *CategoryHandler) ListAllCategoriesWithoutPagination(context *gin.Context) {
	list := Category.FindAllCategories()
	var nameList []string
	for _, name := range list {
		nameList = append(nameList, name.GetName())
	}
	context.JSON(http.StatusOK, gin.H{
		"message":    "Categories listed successfully",
		"categories": nameList,
	})
}
