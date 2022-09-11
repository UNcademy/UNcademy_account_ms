package models

import (
	util "UNcademy_account_ms/utils"
	"gorm.io/gorm"
	"time"
)

type Role string

type User struct {
	UserName       string `gorm:"primaryKey;"`
	UserType       Role   `gorm:"type:ENUM('student', 'professor', 'both' );not null"`
	Password       string `gorm:"type:varchar(255);not null"`
	FullName       string `gorm:"type:varchar(255);not null"`
	Document       int    `gorm:"type:int;unique;not null"`
	Email          string `gorm:"type:varchar(255);unique;not null"`
	Active         bool   `gorm:"type:bool;default:true"`
	DepDocument    string `gorm:"type:varchar(255);not null"`
	CityDocument   string `gorm:"type:varchar(255);not null"`
	Genre          string `gorm:"type:varchar(255);not null"`
	UNMail         string `gorm:"type:varchar(255);unique;not null"`
	Cel            int    `gorm:"type:int;not null"`
	Tel            int    `gorm:"type:int;not null"`
	Age            int    `gorm:"type:int"`
	BirthPlace     string `gorm:"type:varchar(255);not null"`
	Country        string `gorm:"type:varchar(255);not null"`
	BloodType      string `gorm:"type:varchar(100);not null"`
	Address        string `gorm:"type:varchar(255);not null"`
	ArmyCard       bool   `gorm:"type:bool;default:false"`
	MotherFullName string `gorm:"type:varchar(255)"`
	MotherDocument int    `gorm:"type:int"`
	FatherFullName string `gorm:"type:varchar(255)"`
	FatherDocument int    `gorm:"type:int"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.Active = true
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
