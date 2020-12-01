package models

import (
	"gorm.io/gorm"
)

type Cursos struct {
    gorm.Model
	Name 		string `json:"name"`
	Period	 	string `json:"period"`
	Note		string `json:"note"`
	State		string `json:"state"`
}