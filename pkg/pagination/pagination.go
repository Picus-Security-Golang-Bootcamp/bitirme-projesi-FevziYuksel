package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 20
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 100
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "pageSize"
)

// Pages represents a paginated list of data items.
type Pages struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The pageSize parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, pageSize, allProducts int) *Pages {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	pageCount := allProducts / pageSize
	if allProducts%pageSize > 0 {
		pageCount++
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: allProducts,
		PageCount:  pageCount,
	}
}

func NewFromGinRequest(g *gin.Context, allProducts int) *Pages {
	page := parseInt(g.Query(PageVar), 1)
	pageSize := parseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, allProducts)
}

func GetPaginationParametersFromRequest(g *gin.Context) (pageIndex, pageSize int) {
	pageIndex = parseInt(g.Query(PageVar), 1)
	pageSize = parseInt(g.Query(PageSizeVar), DefaultPageSize)
	return pageIndex, pageSize
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil && result > 0 {
		return result
	}
	return defaultValue
}
