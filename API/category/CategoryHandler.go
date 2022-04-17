package CategoryHandler

import (
	Category "FinalProjectGO/Models/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
}

//Constructor
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

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

	categoryList, err := Category.CSVtoCategory("CSV/" + file.Filename)

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

		context.JSON(http.StatusOK, gin.H{
			"message": "Category created successfully",
		})
	}
}
func (c *CategoryHandler) ListAllCategories(context *gin.Context) {
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
