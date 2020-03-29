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
