package main

import (
	"flow/api"

	_ "flow/docs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	
	r.POST("/flow/list", api.FlowList)
	r.POST("/flow/telcheck", api.FlowTelcheck)
	r.POST("/flow/recharge", api.FlowRecharge)
	r.POST("/flow/ordersta", api.FlowOrdersta)

	r.Run()
}
