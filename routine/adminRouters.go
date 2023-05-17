package routine

import (
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	group := r.Group("/admin")
	{
		group.GET("/users", func(c *gin.Context) {
			c.String(200,"/admin/users")
		})

		group.GET("/add", func(c *gin.Context) {
			c.String(200,"/admin/add")
		})
	}

}
