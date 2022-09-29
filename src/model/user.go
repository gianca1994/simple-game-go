package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model            `json:"-"`
	Id                    int    `gorm:"primary_key" json:"id"`
	Username              string `gorm:"type:varchar(100);unique_index" json:"username"`
	Password              string `gorm:"type:varchar(255)" json:"password"`
	Role                  string `gorm:"default:standard" json:"role"`
	Level                 int    `gorm:"default:1" json:"level"`
	Experience            int    `gorm:"default:0" json:"experience"`
	ExperienceToNextLevel int    `gorm:"default:100" json:"experience_to_next_level"`
	DamageMax             int    `gorm:"default:20" json:"damage_max"`
	DamageMin             int    `gorm:"default:5" json:"damage_min"`
	Health                int    `gorm:"default:100" json:"health"`
	DefenseMax            int    `gorm:"default:10" json:"defense_max"`
	DefenseMin            int    `gorm:"default:5" json:"defense_min"`
}

type UserRegisterDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	Username              string `json:"username"`
	Role                  string `json:"role"`
	Level                 int    `json:"level"`
	Experience            int    `json:"experience"`
	ExperienceToNextLevel int    `json:"experience_to_next_level"`
	DamageMax             int    `json:"damage_max"`
	DamageMin             int    `json:"damage_min"`
	Health                int    `json:"health"`
	DefenseMax            int    `json:"defense_max"`
	DefenseMin            int    `json:"defense_min"`
}

func UserToUserResponseDTO(user User) UserResponseDTO {
	var userResponseDTO UserResponseDTO

	userResponseDTO.Username = user.Username
	userResponseDTO.Role = user.Role
	userResponseDTO.Level = user.Level
	userResponseDTO.Experience = user.Experience
	userResponseDTO.ExperienceToNextLevel = user.ExperienceToNextLevel
	userResponseDTO.DamageMax = user.DamageMax
	userResponseDTO.DamageMin = user.DamageMin
	userResponseDTO.Health = user.Health
	userResponseDTO.DefenseMax = user.DefenseMax
	userResponseDTO.DefenseMin = user.DefenseMin

	return userResponseDTO
}
