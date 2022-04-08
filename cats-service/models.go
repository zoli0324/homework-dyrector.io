package main

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name string `json:"cat" gorm:"unique"`
	Art  string `json:"art" gorm:"type:varchar(4096)"`
}
