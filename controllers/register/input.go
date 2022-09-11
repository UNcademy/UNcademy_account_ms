package register

import model "UNcademy_account_ms/models"

type InputRegister struct {
	UserName       string     `json:"user_name" validate:"required,lowercase"`
	UserType       model.Role `json:"user_type" validate:"required"`
	Password       string     `json:"password" validate:"required,gte=8"`
	FullName       string     `json:"full_name" validate:"required,lowercase"`
	Document       int        `json:"document" validate:"required,numeric"`
	DepDocument    string     `json:"dep_document" validate:"required,lowercase"`
	CityDocument   string     `json:"city_document" validate:"required,lowercase"`
	Genre          string     `json:"genre" validate:"required,lowercase"`
	Email          string     `json:"email" validate:"required,email"`
	UNMail         string     `json:"un_mail" validate:"required, regexp=^[a-zA-Z0-9]*@unal.edu.co"`
	BirthPlace     string     `json:"birth_place" validate:"required,lowercase"`
	Cel            int        `json:"cel" validate:"required,numeric"`
	Tel            int        `json:"tel" validate:"numeric"`
	Age            int        `json:"age" validate:"required,numeric"`
	Country        string     `json:"country" validate:"required,lowercase"`
	BloodType      string     `json:"blood_type" validate:"required,lowercase"`
	Address        string     `json:"address" validate:"required"`
	ArmyCard       bool       `json:"army_card" validate:"required,boolean"`
	MotherFullName string     `json:"mother_full_name" validate:"lowercase"`
	MotherDocument int        `json:"mother_document" validate:"numeric"`
	FatherFullName string     `json:"father_full_name" validate:"lowercase"`
	FatherDocument int        `json:"father_document" validate:"numeric"`
}
