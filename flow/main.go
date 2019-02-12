package main

import (
	"flow/api"

	_ "flow/docs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	
	r.POST("/flow/list", api.Flowlist)

	r.Run()
}
