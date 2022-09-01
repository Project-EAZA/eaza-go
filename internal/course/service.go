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
	FindCourseById(uuid string, result *model.Course) error
}

type ServiceImpl struct {
	DB *database.DataBase
}

func (s *ServiceImpl) FindCourse(name string, number int, result *model.Course) error {
	err := s.DB.Courses().FindOne(context.TODO(), bson.M{"name": name, "courseNumber": number}).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) FindCourseById(uuid string, result *model.Course) error {
	opts := options.FindOne().SetProjection(bson.D{{"description", 0}, {"requirement", 0}})
	err := s.DB.Courses().FindOne(context.TODO(), bson.M{"uuid": uuid}, opts).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
