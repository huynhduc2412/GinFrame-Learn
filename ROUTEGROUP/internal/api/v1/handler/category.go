package v1handler

import (
	"RouteGroup/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var validCategory = map[string]struct{} {
	"java" : struct{}{},
	"python" : struct{}{},
	"golang" : struct{}{},
}

func NewCategorytHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")
	if err := utils.ValidationInList("Category" , category , validCategory) ; err != nil {
		ctx.JSON(http.StatusBadRequest , gin.H{
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{
		"message" : "Category found",
		"category" : category,
	})
}