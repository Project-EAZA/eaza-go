package course

import "eaza-go/internal/course/model"

type Overview struct {
	Title         string          `json:"title"`
	CatalogNumber string          `json:"catalogNumber"`
	Subject       model.Subject   `json:"subject"`
	GeneralEd     model.GeneralEd `json:"generalEd"`
	Level         []model.Level   `json:"level"`
	Breadths      []model.Breadth `json:"breadths"`
	Repeatable    string          `json:"repeatable"`
	GPA           float32         `json:"GPA"`
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
		GPA:           c.AvgGPA(),
	}
}
