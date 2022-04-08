package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateCatInput struct {
	Name string `json:"cat"`
	Art  string `json:"art"`
}

type UpdateCatInput struct {
	Name string `json:"cat"`
	Art  string `json:"art"`
}

// GET /cats
// Get all cats
func FindCats(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cats []Cat
	db.Find(&cats)

	c.JSON(http.StatusOK, gin.H{"data": cats})
}

// POST /cats
// Create new cat
func CreateCat(c *gin.Context) {
	// Validate input
	var input CreateCatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create cat
	cat := Cat{Name: input.Name, Art: input.Art}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&cat)

	c.JSON(http.StatusOK, gin.H{"data": cat})
}

// GET /cats/:id
// Find a cat
func FindCat(c *gin.Context) { // Get model if exist
	var cat Cat

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cat})
}

// PATCH /cats/:id
// Update a cat
func UpdateCat(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var cat Cat
	if err := db.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&cat).Updates(Cat{Name: input.Name, Art: input.Art})

	c.JSON(http.StatusOK, gin.H{"data": cat})
}

// DELETE /cats/:id
// Delete a cat
func DeleteCat(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book Cat
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func GenerateCat(c *gin.Context) {
	var count int64
	var cat Cat

	db := c.MustGet("db").(*gorm.DB)

	if c.Query("name") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is missing"})
		return
	}

	db.Model(&Cat{}).Count(&count)

	name := c.Query("name")

	sum := 0
	for i := range name {
		sum += int(name[i])
	}

	index := 0

	if count != 0 {
		index = (sum % int(count-1)) + 1
	}

	result := db.Find(&cat, index)

	if result.Error != nil {
		log.Println("db error: ", result.Error.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cat})
}
