package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model            `json:"-"`
	Id                    int    `gorm:"primary_key" json:"id"`
	Username              string `gorm:"type:varchar(100);unique_index" json:"username"`
	Password              string `gorm:"type:varchar(255)" json:"password"`
	Level                 int    `gorm:"not null" json:"level"`
	Experience            int    `gorm:"not null" json:"experience"`
	ExperienceToNextLevel int    `gorm:"not null" json:"experience_to_next_level"`
	DamageMax             int    `gorm:"not null" json:"damage_max"`
	DamageMin             int    `gorm:"not null" json:"damage_min"`
	Health                int    `gorm:"not null" json:"health"`
	DefenseMax            int    `gorm:"not null" json:"defense_max"`
	DefenseMin            int    `gorm:"not null" json:"defense_min"`
}

