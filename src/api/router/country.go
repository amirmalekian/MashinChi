package router

import (
	"github.com/amirmalekian/Mashinchi/src/api/handler"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/gin-gonic/gin"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
}
