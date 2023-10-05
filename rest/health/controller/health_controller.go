package controller

import (
	"github.com/gin-gonic/gin"
)

func GetHealthCheck(c *gin.Context) {
	c.String(200, "ok")
}