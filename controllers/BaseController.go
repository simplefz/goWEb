package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//todo 数据验证 方法验证 用户验证
func IndexApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test" :"67890",
	})
}