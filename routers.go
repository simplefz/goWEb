package main

/*import (
	"github.com/fanpengjie/project/common"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/fanpengjie/project/common/auth"
	"net/http"
)

func InitRouter() *gin.Engine {
		router:=gin.New()

		if viper.GetString("debug") == "debug"{   			//性能分析工具
			pprof.Register(router)
		}

		router.Use(gin.Logger())
		router.Use(common.RegisterSession())

		router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth"))

		router.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "找不到该路由",
			})
			return
		})

		router.NoMethod(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "找不到该方法",
			})
			return
		})


	return  router
}*/