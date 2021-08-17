package users

import (
	"net/http"
	"strconv"

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
	userId, errP := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if errP != nil {
		err := errors.NewBadRequestErr("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}
	user, errG := services.GetUser(userId)
	if errG != nil {
		ctx.JSON(errG.Status, errG)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
