package utils

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

func SonicJSON(c *gin.Context, code int, obj interface{}) {
	jsonBytes, err := sonic.Marshal(obj)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "JSON encode error"})
		return
	}
	c.Data(code, "application/json; charset=utf-8", jsonBytes)
}
