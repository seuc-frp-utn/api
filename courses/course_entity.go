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
	Hours uint64 `json:"hours"`
	Price uint64 `json:"price"`
	Classes uint64 `json:"classes"`
	Link string `json:"link"`
	Teacher string `json:"teacher"`
}

type CourseCreate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Started bool `json:"started"`
	Finished bool `json:"finished"`
	Hours uint64 `json:"hours"`
	Price uint64 `json:"price"`
	Classes uint64 `json:"classes"`
	Link string `json:"link"`
	Teacher string `json:"teacher"`
}

type CourseUpdate struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Image *string `json:"image,omitempty"`
	Started *bool `json:"started,omitempty"`
	Finished *bool `json:"finished,omitempty"`
	Hours *uint64 `json:"hours,omitempty"`
	Price *uint64 `json:"price,omitempty"`
	Classes *uint64 `json:"classes,omitempty"`
	Link *string `json:"link,omitempty"`
	Teacher *string `json:"teacher,omitempty"`
}
