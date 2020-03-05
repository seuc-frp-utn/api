package courses

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Started bool `json:"started"`
	Finished bool `json:"finished"`
	Hours uint `json:"hours"`
	Price uint `json:"price"`
	Months uint `json:"months"`
	Link string `json:"link"`
	Teacher string `json:"teacher"`
}

type CourseCreate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Started bool `json:"started"`
	Finished bool `json:"finished"`
	Hours uint `json:"hours"`
	Price uint `json:"price"`
	Months uint `json:"months"`
	Link string `json:"link"`
	Teacher string `json:"teacher"`
}

type CourseUpdate struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Image *string `json:"image,omitempty"`
	Started *bool `json:"started,omitempty"`
	Finished *bool `json:"finished,omitempty"`
	Hours *uint `json:"hours,omitempty"`
	Price *uint `json:"price,omitempty"`
	Months *uint `json:"months,omitempty"`
	Link *string `json:"link,omitempty"`
	Teacher *string `json:"teacher,omitempty"`
}
