package course

import "eaza-go/internal/course/model"

type Overview struct {
	CourseId      string              `json:"courseId"`
	Title         string              `json:"title"`
	CatalogNumber string              `json:"catalogNumber"`
	Subject       model.Subject       `json:"subject"`
	GeneralEd     model.GeneralEd     `json:"generalEd"`
	Level         []model.Level       `json:"level"`
	Breadths      []model.Breadth     `json:"breadths"`
	EthnicStudies model.EthnicStudies `json:"ethnicStudies"`
	Repeatable    string              `json:"repeatable"`
	GPA           float32             `json:"GPA"`
}

type GradesDistribution struct {
	Teachings []model.Teaching `json:"teachings"`
}

func NewOverviewFromCourse(c model.Course) Overview {
	return Overview{
		Title:         c.Title,
		CatalogNumber: c.CatalogNumber,
		Subject:       c.Subject,
		GeneralEd:     c.GeneralEd,
		Level:         c.Level,
		Breadths:      c.Breadths,
		Repeatable:    c.Repeatable,
		EthnicStudies: c.EthnicStudies,
		GPA:           c.AvgGPA(),
	}
}
