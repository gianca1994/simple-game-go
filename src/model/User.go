package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id         int    `gorm:"primary_key" json:"id"`
	Name       string `gorm:"size:255;not null;unique" json:"name"`
	DamageMax  int    `gorm:"not null" json:"damage_max"`
	DamageMin  int    `gorm:"not null" json:"damage_min"`
	Health     int    `gorm:"not null" json:"health"`
	DefenseMax int    `gorm:"not null" json:"defense_max"`
	DefenseMin int    `gorm:"not null" json:"defense_min"`
}
