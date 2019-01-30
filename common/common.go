package common

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterSession() gin.HandlerFunc  {
	store,_:=sessions.NewRedisStore(
		10,
		"tcp",
		viper.GetString("cache.host")+":"+viper.GetString("cache.port"),
		viper.GetString("cache.password"),
		[]byte("secret"))
	return sessions.Sessions("shangce",store)
}