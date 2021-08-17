package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1998/bookstore-users-api/domain/users"
	"github.com/kartik1998/bookstore-users-api/services"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	bytes, errR := ioutil.ReadAll(ctx.Request.Body)
	if errR != nil {
		// handle err
		fmt.Println("err ", errR)
		return
	}
	errP := json.Unmarshal(bytes, &user)
	if errP != nil {
		// handle err
		fmt.Println("errP ", errP)
		return
	}
	result, errS := services.CreateUser(user)
	if errS != nil {
		// handle user creation err
		return
	}
	fmt.Println("result ", result)
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Not Implemented")
}
