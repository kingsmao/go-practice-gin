package routine

import "github.com/gin-gonic/gin"

func OpenApiRouters(r *gin.Engine) {
	group := r.Group("/openapi")
	{
		group.GET("/health", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"message":"ok",
			})
		})

		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"message":"pong",
			})
		})
	}
}
