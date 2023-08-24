package models

import "gorm.io/gorm"

type riddle struct {
	gorm.Model
	Riddle    string `json:"question" gorm:"text;not null;default null"`
	Solution  string `json:"solution" gorm:"text;not null;default null"`
	Published string `json: "published" gorm:"text;not null;default null"`
}
