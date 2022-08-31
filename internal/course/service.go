package course

import (
	"context"
	"eaza-go/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	FindCourse(name string, number int, result *Course) error
	FindCourseById(uuid string, result *Course) error
}

type ServiceImpl struct {
	DB *database.DataBase
}

func (s *ServiceImpl) FindCourse(name string, number int, result *Course) error {
	err := s.DB.Courses().FindOne(context.TODO(), bson.M{"name": name, "courseNumber": number}).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) FindCourseById(uuid string, result *Course) error {
	err := s.DB.Courses().FindOne(context.TODO(), bson.M{"uuid": uuid}).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
