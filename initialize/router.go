package initialize

import (
	"net/http"

	"github.com/akazwz/imgin/middleware"
	"github.com/akazwz/imgin/model/response"
	"github.com/akazwz/imgin/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers() *gin.Engine {
	r := gin.Default()

	/* cors 使用跨域中间件 */
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	/* 404 路由不存在 */
	r.NoRoute(response.NotFound)

	//Teapot  418
	r.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	/* 公共路由 */
	publicRouterV1 := r.Group("v1")
	{
		router.InitBaseRouter(publicRouterV1)
		router.InitCloudRouter(publicRouterV1)
	}

	privateGroupV1 := r.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(privateGroupV1)
		router.InitImageRouter(privateGroupV1)
	}

	return r
}
