package v1handler

import (
	"RouteGroup/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

type GetUsersByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`

}

type GetUsersByUUidV1Param struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all users (V1)",
	})
}

func (u *UserHandler) GetUsersByIdV1(ctx *gin.Context) {
	var params GetUsersByIdV1Param
	if err := ctx.ShouldBindUri(&params) ; err != nil {
		ctx.JSON(http.StatusBadRequest , utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by id (V1)",
		"user_id" : params.ID,
	})
}

func (u *UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	var params GetUsersByUUidV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest , utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by uuid (V1)",
		"user_id" : params.Uuid,
	})
}

func (u *UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user (V1)",
	})
}

func (u *UserHandler) PutUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user (V1)",
	})
}

func (u *UserHandler) DeleteUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete user (V1)",
	})
}