package course

import (
	"eaza-go/internal/common"
	"eaza-go/internal/course/model"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetCourseByNumberAndName(ctx *fiber.Ctx) error
	GetCourseByAbbrAndNumber(ctx *fiber.Ctx) error
}

type ControllerImpl struct {
	service Service
}

func (c *ControllerImpl) GetCourseByAbbrAndNumber(ctx *fiber.Ctx) error {
	req := new(AbbrAndNumberRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	var course model.Course
	err := c.service.FindCourseBySubject(req.Abbr, req.Number, &course)
	if err != nil {
		return err
	}

	err = ctx.JSON(common.NewSuccessResponse(NewOverviewFromCourse(course)))

	if err != nil {
		return err
	}
	return nil
}

func NewController(service Service) Controller {
	return &ControllerImpl{service: service}
}

func (c *ControllerImpl) GetCourseByNumberAndName(ctx *fiber.Ctx) error {
	req := new(NameAndNumberRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	var course model.Course
	err := c.service.FindCourse(req.Name, req.Number, &course)
	if err != nil {
		return err
	}

	err = ctx.JSON(common.NewSuccessResponse(course))

	if err != nil {
		return err
	}

	return nil
}
