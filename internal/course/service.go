package course

import (
	"context"
	"eaza-go/database"
	"eaza-go/internal/course/model"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	FindCourse(name string, number int, result *model.Course) error
	FindCourseBySubject(subAbbr string, number string, result *model.Course) error
	FindCourseById(uuid string, result *model.Course) error
	FindGradesBySubject(subAbbr string, number string, result *GradesDistribution) error
}

func NewService() Service {
	return &ServiceImpl{DB: database.DB}
}

type ServiceImpl struct {
	DB *database.DataBase
}

func (s *ServiceImpl) FindGradesBySubject(subAbbr string, number string, result *GradesDistribution) error {
	// Try to get course from redis cache
	var c model.Course
	err := getCache(courseKey(subAbbr, number), &c)
	if err == nil {
		result.Teachings = c.Teachings
		return nil
	}

	opts := options.FindOne().SetProjection(bson.D{
		{"teachings", 1},
	})

	return s.DB.Course().FindOne(context.TODO(), bson.M{
		"catalogNumber":            number,
		"subject.shortDescription": subAbbr,
	}, opts).Decode(result)
}

func (s *ServiceImpl) FindCourseBySubject(subAbbr string, number string, result *model.Course) error {
	key := courseKey(subAbbr, number)
	if err := getCache(key, result); err == nil {
		return nil
	}

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

	err := s.DB.Course().FindOne(context.TODO(), bson.M{
		"catalogNumber":            number,
		"subject.shortDescription": subAbbr,
	}, opts).Decode(result)
	if err != nil {
		return err
	}

	return writeCache(result)
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

func getCache(key string, result *model.Course) error {
	// Try to get course from redis cache
	bytes, err := database.RedisClient.Get(key)
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}
	return nil
}

func writeCache(course *model.Course) error {
	key := courseKey(course.Subject.ShortDescription, course.CatalogNumber)
	bytes, err := msgpack.Marshal(course)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = database.RedisClient.DefaultSet(key, bytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func courseKey(subAbbr string, number string) string {
	return subAbbr + number
}
