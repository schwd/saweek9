package controller

import (
	"github.com/schwd/sa-64-example/backend/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /doctors
// List all doctors
func ListDoctors(c *gin.Context) {
	var doctors []entity.Doctor
	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

// GET /doctor/:id
//Get doctor by id
func GetDoctor(c *gin.Context) {
	var doctor entity.Doctor
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM doctors WHERE id = ?", id).Scan(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// POST / doctors
func CreateDoctor(c *gin.Context) {
	var doctor entity.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}


// PATCH /doctors
func UpdateDoctor(c *gin.Context) {
	var doctor entity.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", doctor.ID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}
	if err := entity.DB().Save(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// DELETE / doctor/id
func DeleteDoctor(c *gin.Context){
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM doctors WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

