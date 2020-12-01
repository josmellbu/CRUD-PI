package models

import (
	"gorm.io/gorm"
)

type Estudiante struct {
    gorm.Model
	Name 		string `json:"name"`
	Paternal 	string `json:"paternal"`
	Maternal	string `json:"maternal"`
	Age  		string `json:"age"`
	State		string `json:"state"`
}