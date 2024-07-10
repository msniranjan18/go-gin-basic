package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Param1 will mapped to param1
// output will have the json tag name
type URI struct {
	Param1 string `json:"param_value1" uri:"param1"`
	Param2 string `json:"param_value2" uri:"param2"`
}

// without json tag output will have the struct fields name
// type URI struct {
// 	Param1 string `uri:"param1"`
// 	Param2 string `uri:"param2"`
// }

func main() {
	fmt.Println("use-of-uri-struct-tag-with-gin-new-engine")
	engine := gin.New()
	// adding path params to router
	engine.GET("/test/:param1/test/:param2", func(context *gin.Context) {
		uri := URI{}
		// binding to URI
		if err := context.BindUri(&uri); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusAccepted, &uri)
	})
	engine.Run(":3000")
}

/*
Output:
HTTP/1.1 202 Accepted
Content-Type: application/json; charset=utf-8
Date: Wed, 10 Jul 2024 12:10:59 GMT
Content-Length: 29
Connection: close

{
  "Param1": "50",
  "Param2": "20"
}
*/

/*
Note: https://github.com/msniranjan18/go-excercises/edit/master/go-gin/gin-binding-poc.go same code.
*/
