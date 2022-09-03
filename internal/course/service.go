package course

import (
	"context"
	"eaza-go/database"
	"eaza-go/internal/course/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	FindCourse(name string, number int, result *model.Course) error
	FindCourseBySubject(subAbbr string, number string, result *model.Course) error
	FindCourseById(uuid string, result *model.Course) error
}

type ServiceImpl struct {
	DB *database.DataBase
}

func (s *ServiceImpl) FindCourseBySubject(subAbbr string, number string, result *model.Course) error {
	opts := options.FindOne().SetProjection(bson.D{
		{"breadths", 1},
		{"levels", 1},
		{"generalEd", 1},
		{"ethnicStudies", 1},
		{"teachings", 1},
		{"title", 1},
		{"catalogNumber", 1},
		{"subject", 1},
		{"repeatable", 1},
	})

	return s.DB.Course().FindOne(context.TODO(), bson.M{
		"catalogNumber":            number,
		"subject.shortDescription": subAbbr,
	}, opts).Decode(result)
}

func (s *ServiceImpl) FindCourse(name string, number int, result *model.Course) error {
	err := s.DB.Course().FindOne(context.TODO(), bson.M{"name": name, "courseNumber": number}).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) FindCourseById(uuid string, result *model.Course) error {
	opts := options.FindOne().SetProjection(bson.D{{"description", 0}, {"requirement", 0}})
	err := s.DB.Course().FindOne(context.TODO(), bson.M{"uuid": uuid}, opts).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
