package router

import (
	"bewell-backend-challenge/internal/adapter/handler"
	"bewell-backend-challenge/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, c *config.Config) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handler := handler.New()

	v1 := r.Group("/v1")
	v1.POST("/clean-orders", handler.CleanOrders)

	return r
}
