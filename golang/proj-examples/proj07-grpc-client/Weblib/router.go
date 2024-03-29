package Weblib

import (
	"proj07-grpc-client/Services"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(prodService Services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(initMiddleware(prodService)) // 使用context进行传值的示例

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
	}

	return ginRouter
}
