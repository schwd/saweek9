package controller

import (
	"github.com/schwd/sa-64-example/backend/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /watch_videos
func CreateRefer(c *gin.Context) {

	var MedicalHistory entity.MedicalHistory
	var medicalrecord entity.MedicalRecord
	var refer entity.Refer
	var doctor entity.Doctor
	var hospital entity.Hospital

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร refer
	if err := c.ShouldBindJSON(&refer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id medrec
	if tx := entity.DB().Where("id = ?", refer.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id doctor
	if tx := entity.DB().Where("id = ?", refer.MedicalRecordID).First(&medicalrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id disease
	if tx := entity.DB().Where("id = ?", refer.HospitalID).First(&hospital); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disease not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id department
	if tx := entity.DB().Where("id = ?", refer.MedicalHistoryID).First(&MedicalHistory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalHistory not found"})
		return
	}

	// 12: สร้าง WatchVideo
	wv := entity.Refer{
		MedicalRecord: medicalrecord,            // โยงความสัมพันธ์กับ Entity medrec
		Doctor:        doctor,                   // โยงความสัมพันธ์กับ Entity Doctor
		MedicalHistory: MedicalHistory,                   // โยงความสัมพันธ์กับ Entity Doctor
		Hospital:		hospital,
		Cause:     		refer.Cause, // ตั้งค่าฟิลด์ diag
		Date:          MedicalHistory.Date,      // ตั้งค่าฟิลด์ date
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /watchvideo/:id
func GetRefer(c *gin.Context) {
	var refer entity.Refer
	id := c.Param("id")
	if err := entity.DB().Preload("MedicalRecord").Preload("Doctor").Preload("Hospital").Preload("MedicalHistory").Raw("SELECT * FROM refers WHERE id = ?", id).Find(&refer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": refer})
}

// GET /watch_videos
func ListRefer(c *gin.Context) {
	var refers []entity.Refer
	if err := entity.DB().Preload("MedicalRecord").Preload("Doctor").Preload("Hospital").Preload("MedicalHistory").Raw("SELECT * FROM refers").Find(&refers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": refers})
}

// DELETE /watch_videos/:id
func DeleteRefer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM refers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalHistory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateRefer(c *gin.Context) {
	var refer entity.MedicalHistory
	if err := c.ShouldBindJSON(&refer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", refer.ID).First(&refer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalHistory not found"})
		return
	}

	if err := entity.DB().Save(&refer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": refer})
}