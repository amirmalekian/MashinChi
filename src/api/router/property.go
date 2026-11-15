package router

import (
	"github.com/amirmalekian/Mashinchi/src/api/handler"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
