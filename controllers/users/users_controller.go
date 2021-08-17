package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1998/bookstore-users-api/domain/users"
	"github.com/kartik1998/bookstore-users-api/services"
	"github.com/kartik1998/bookstore-users-api/utils/errors"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Not Implemented")
}
