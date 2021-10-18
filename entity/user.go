package entity

import (
    "gorm.io/gorm"
    "time"
)

type Doctor struct {
	gorm.Model
	Name		string
	Tel			string
	Email		string			`gorm:"uniqueIndex"`
	Password	string
	// 1 Doctor เป็นเจ้าของได้หลาย MedicalHistory
	MedicalHistorys		[]MedicalHistory	`gorm:"foreignKey:DoctorID"`
	// 1 Doctor เป็นเจ้าของได้หลาย Refer
	Refers				[]Refer				`gorm:"foreignKey:DoctorID"`
}

type MedicalHistory struct {
	gorm.Model
	Diagnosis			string
	Treatment			string
	DiseaseID			int
	Date				time.Time
	// DoctorID ทำหน้าที่เป็น FK
	DoctorID			*uint
	// เป็นข้อมูล Doctor ใช้เพื่อให้ join ตาราง
	Doctor				Doctor

	// MedRecID ทำหน้าที่เป็น FK
	MedicalRecordID			*uint
	// เป็นข้อมูล MedicalRecord ใช้เพื่อให้ join ตาราง
	MedicalRecord		MedicalRecord

	// 1 MedicalHistory เป็นเจ้าของได้หลาย Refer
	Refers				[]Refer				`gorm:"foreignKey:MedicalHistoryID"`
	DepartmentID		int
}

type Refer struct {
	gorm.Model
	Date				time.Time
	Cause				string
	// DoctorID ทำหน้าที่เป็น FK
	DoctorID			*uint
	// เป็นข้อมูล Doctor ใช้เพื่อให้ join ตาราง
	Doctor				Doctor

	// MedHisID ทำหน้าที่เป็น FK
	MedicalHistoryID			*uint
	// เป็นข้อมูล MedicalHistory ใช้เพื่อให้ join ตาราง
	MedicalHistory				MedicalHistory

	// MedRecID ทำหน้าที่เป็น FK
	MedicalRecordID			*uint
	// เป็นข้อมูล MedicalRecord ใช้เพื่อให้ join ตาราง
	MedicalRecord				MedicalRecord

	// HospitalID ทำหน้าที่เป็น FK
	HospitalID			*uint
	// เป็นข้อมูล Hospital ใช้เพื่อให้ join ตาราง
	Hospital			Hospital
}

type MedicalRecord struct {
	gorm.Model
	HospitalNumber			string
	PersonalID				string		`gorm:"uniqueIndex"`
	NameTitleID				int
	PatientName				string
	PatientAge				*uint
	PatientDob				time.Time
	PatientTel				string
	HealthInsuranceID		int
	RegisterDate			time.Time
	// 1 MedicalRecord เป็นเจ้าของได้หลาย MedicalHistory
	MedicalHistorys				[]MedicalHistory		`gorm:"foreignKey:MedicalRecordID"`

	// 1 MedicalRecord เป็นเจ้าของได้หลาย Refer
	Refers				[]Refer							`gorm:"foreignKey:MedicalRecordID"`
}

type Hospital struct {
	gorm.Model
	Name		string
	Tel			string
	// 1 Hospital เป็นเจ้าของได้หลาย Refer
	Refers		[]Refer		`gorm:"foreignKey:HospitalID"`
}