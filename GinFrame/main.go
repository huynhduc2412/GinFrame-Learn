package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/demo" , func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK , gin.H{"message" : "Hello World"})
	})
	r.GET("/users" , func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK , gin.H{"data" : "List Users"})
	})
	r.GET("/user/:user_id" , func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		ctx.JSON(http.StatusOK , gin.H{
			"data" : "Info User",
			"user_id" : user_id,
		})
	})
	r.GET("/products" , func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK , gin.H{"data" : "List Products"})
	})
	r.GET("/product/:product_name" , func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")

		price := ctx.Query("price")

		ctx.JSON(http.StatusOK , gin.H{
			"data" : "Info Product",
			"product_name" : product_name,
			"product_price" : price,
		})
	})

	r.Run()
	// r.Run(":8080")
}