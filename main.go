package main

import (
	"flag"
	"github.com/fanpengjie/project/common"
	"github.com/fanpengjie/project/common/auth"
	"github.com/fanpengjie/project/config"
	"github.com/fanpengjie/project/routes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var (
	cfg = flag.String("config","","api service config  path")
)

func main() {
	err:=config.Init(*cfg)
	if err != nil {
		println(err.Error())
	}
	gin.SetMode(viper.GetString("debug"))
	router := InitRouter()
	router.Run(":8070")

}


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

	routes.RegisterApiRouter(router)
	return  router
}



