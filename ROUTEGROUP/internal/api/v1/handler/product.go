package v1handler

import (
	"RouteGroup/utils"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)
var (
	slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
)


type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (p *ProductHandler)  GetProductsV1(ctx *gin.Context) {
	search := ctx.Query("search")
	if err := utils.ValidationRequired("Search" , search) ; err != nil {
		ctx.JSON(http.StatusBadRequest , gin.H{
			"error" : err.Error(),
		})
		return
	}
	if err := utils.ValidationStringLength("Search" , search , 3 , 50) ; err != nil {
		ctx.JSON(http.StatusBadRequest , gin.H{
			"error" : err.Error(),
		})
		return
	}

	if err := utils.ValidationRegex("Search" , search , searchRegex , "must contain only letters , numbers and spaces") ; err != nil{
		ctx.JSON(http.StatusBadRequest , gin.H{
			"errors" : err.Error(),
		})
		return
	}
	limitStr := ctx.DefaultQuery("limit" , "10")
	limit , err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest , gin.H{
			"error" : "Limit must be a positive number",
		})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{
			"message" : "List all Products (V1)",
			"search" : search,
			"linit" : limit,
	})
}

func (p *ProductHandler)  GetProductsBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := utils.ValidationRegex("Slug" , slug , slugRegex , "must contain only lowercase letters , numbers , hyphens and dots") ; err != nil{
		ctx.JSON(http.StatusBadRequest , gin.H{
			"errors" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{
			"message" : "Get Product by Slug (V1)",
			"slug" : slug,
	})
}

func (p *ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated , gin.H{
			"message" : "Create Product (V1)",
		})
}

func (p *ProductHandler)  PutProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{
			"message" : "Update Product (V1)",
		})
}

func (p *ProductHandler)  DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent , gin.H{
			"message" : "Delete Product (V1)",
		})
}
