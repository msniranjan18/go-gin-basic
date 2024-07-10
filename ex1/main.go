package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Field name shall be Upper case for json tag to work
// binding only  works when the fields are in Upper case (exportable)
type user struct {
	FName string `binding:"required" json:"f_name,omitempty"`
	LName string `binding:"required" json:"l_name,omitempty"`
}

// binding works when the fields are in Upper case (exportable)
// type user struct {
// 	FName string `binding:"required"` // Field name shall be Upper case as useing json tag
// 	LName string `binding:"required"` // Field name shall be Upper case as useing json tag
// }

var (
	users = []user{}
)

func main() {
	fmt.Println("Go-Gin-basic-use")
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "welcome!!"})
	})

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "users List", "users": users})
	})

	server.POST("/user", func(ctx *gin.Context) {
		u := user{}
		err := ctx.ShouldBindJSON(&u)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		users = append(users, u)
		fmt.Println(u)
		ctx.JSON(http.StatusCreated, gin.H{"message": "user created!", "user": u})
	})

	server.Run(":8080")
}
