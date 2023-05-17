package main

import (
	"github.com/gin-gonic/gin"
	"go-practice-gin/routine"
)


func main() {
	//创建默认路由引擎
	r := gin.Default()
	//配置路由

	routine.AdminRouter(r)
	
	routine.OpenApiRouters(r)

	//启动web服务
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}