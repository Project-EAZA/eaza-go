package course

import (
	"eaza-go/internal/course/model"
)

type BriefCourse struct {
	Name         string          `json:"name"`
	CourseNumber int             `json:"courseNumber"`
	Breadths     []model.Breadth `json:"breadths"`
	Level        model.Level     `json:"level"`
	GE           model.GE        `json:"GE"`
	Ethnic       model.Ethnic    `json:"ethnic"`
	GPA          float32         `json:"GPA"`
}

func NewBriefCourse(c *model.Course) BriefCourse {
	return BriefCourse{
		Name:         c.Name,
		CourseNumber: c.CourseNumber,
		Breadths:     c.Breadths,
		Level:        c.Level,
		GE:           c.GE,
		Ethnic:       c.Ethnic,
		GPA:          c.AvgGPA(),
	}
}
