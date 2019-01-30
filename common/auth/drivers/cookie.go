package drivers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"net/http"
)

var store = sessions.NewCookieStore([]byte(viper.GetString("appSecret")))

type cookieAuthManager struct {
	name string
}

func NewCookieAuthDriver () *cookieAuthManager  {
		return &cookieAuthManager{
			name: viper.GetString("sessionName"),
		}
}

func (cookie *cookieAuthManager) Check(c *gin.Context)bool  {
		session ,err := store.Get(c.Request,cookie.name)
		if err!=nil{
			return  false
		}
		if session == nil{
			return  false
		}
		if session.Values == nil{
			return  false
		}
		return  true
}

func (cookie *cookieAuthManager) User (c *gin.Context)interface{}  {
		session,err := store.Get(c.Request,cookie.name)
		if err !=nil{
			return  session.Values
		}
		return session.Values
}

func (cookie *cookieAuthManager) Login(http *http.Request,w http.ResponseWriter,user map[string]interface{})interface{} {
		session, err := store.Get(http,cookie.name)
		if err !=nil{
			return  false
		}
		session.Values["id"] = user["id"]
		session.Save(http,w)
		return  true
}

func (cookie *cookieAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool{
	    session, err := store.Get(http,cookie.name)
		if err !=nil{
			return  false
		}
	   session.Values["id"] = nil
	   session.Save(http, w)
	   return true
}

