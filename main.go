package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	r.GET("/add", func(ctx *gin.Context) {
		a, err := strconv.Atoi(ctx.Query("a"))
		if err != nil {
			ctx.String(400, "Error with 'a' param: " + err.Error())
			return
		}
		b, err := strconv.Atoi(ctx.Query("b"))
		if err != nil {
			ctx.String(400, "Error with 'b' param: " + err.Error())
			return
		}
		ctx.String(200, fmt.Sprintf("a + b = %d", Add(a, b)))
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}