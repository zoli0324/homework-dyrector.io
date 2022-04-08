package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.Use(cors.Default())

	r.GET("/meow", GenerateCat)
	r.GET("/cats", FindCats)
	r.POST("/cats", CreateCat)
	r.GET("/cats/:id", FindCat)
	r.PATCH("/cats/:id", UpdateCat)
	r.DELETE("cats/:id", DeleteCat)
	return r
}
