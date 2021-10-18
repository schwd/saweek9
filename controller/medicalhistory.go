package controller

import (
	"github.com/schwd/sa-64-example/backend/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /medicalhistorys
func CreateMedicalHistory(c *gin.Context) {
	var medicalhistory entity.MedicalHistory
	if err := c.ShouldBindJSON(&medicalhistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&medicalhistory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicalhistory})
}

// GET /medicalhistory/:id
func GetMedicalHistory(c *gin.Context) {
	var medicalhistory entity.MedicalHistory
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medical_histories WHERE id = ?", id).Scan(&medicalhistory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalhistory})
}

// GET /medicalhistorys
func ListMedicalHistories(c *gin.Context) {
	var medicalhistorys []entity.MedicalHistory
	if err := entity.DB().Raw("SELECT * FROM medical_histories").Scan(&medicalhistorys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalhistorys})
}

// DELETE /medicalhistorys/:id
func DeleteMedicalHistory(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medical_histories WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicalhistory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /medicalhistorys
func UpdateMedicalHistory(c *gin.Context) {
	var medicalhistory entity.MedicalHistory
	if err := c.ShouldBindJSON(&medicalhistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicalhistory.ID).First(&medicalhistory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicalhistory not found"})
		return
	}

	if err := entity.DB().Save(&medicalhistory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalhistory})
}