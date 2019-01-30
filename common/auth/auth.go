package auth

import (
	"github.com/fanpengjie/project/common/auth/drivers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var driverList = map[string]func()Auth {
	"cookie": func () Auth {
		return drivers.NewCookieAuthDriver()
	},
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request,w http.ResponseWriter,user map[string]interface{})interface{}
	Logout(http *http.Request,w http.ResponseWriter) bool
}

//生成验证驱动
func GenerateAuthDriver(string string) *Auth {
	var authDriver Auth
	authDriver = driverList[string]()
	return &authDriver
}

//注册
func RegisterGlobalAuthDriver(authkey string , key string) gin.HandlerFunc  {
		return func(c *gin.Context) {
				driver := GenerateAuthDriver(authkey)
				c.Set(key,driver)
				c.Next()
		}
}


func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		if !(*driver).Check(c) {
			c.JSON(http.StatusOK,gin.H{
				"title": "尚未登录，请登录",
			})
			c.Abort()
		}
		c.Next()
	}
}


func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(*Auth)
	return (*authDriver).User(c).(map[string]interface{})
}


