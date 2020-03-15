package diplomas

import "github.com/jinzhu/gorm"

type Diploma struct {
	gorm.Model
	UUID string `json:"uuid"`
	Token string `json:"token"`
	Course string `json:"course"`
	Dean string `json:"dean"`
	Secretary string `json:"secretary"`
	Teacher string `json:"teacher"`
}

type DiplomaCreate struct {
	Course string `json:"course"`
	Dean string `json:"dean"`
	Secretary string `json:"secretary"`
	Teacher string `json:"teacher"`
}

type DiplomaUpdate struct {
	Course *string `json:"course"`
	Dean *string `json:"dean"`
	Secretary *string `json:"secretary"`
	Teacher *string `json:"teacher"`
}