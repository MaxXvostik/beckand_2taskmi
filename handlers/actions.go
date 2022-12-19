package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method":  http.MethodGet,
			"code":    http.StatusOK,
			"message": "Hello word!",
		},
	})
}
