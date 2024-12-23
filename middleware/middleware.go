package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessControllAllowConfig struct {
	Origin  string
	Methods string
	Headers string
}

func AccessControllAllowfunc(config AccessControllAllowConfig) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", config.Methods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", config.Headers)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}
