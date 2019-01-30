package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImportStudent(c *gin.Context)  {
  	projectId :=c.Param("projectId")
  	c.JSON(http.StatusOK,gin.H{
  		"project_id" :projectId,
	})


}