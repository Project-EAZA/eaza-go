package course

import (
	"eaza-go/internal/common"
	"eaza-go/internal/course/model"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetCourseByNumberAndName(ctx *fiber.Ctx) error
	GetCourseByAbbrAndNumber(ctx *fiber.Ctx) error
	GetGradesBySubject(ctx *fiber.Ctx) error
}

func NewController(service Service) Controller {
	return &ControllerImpl{service: service}
}

type ControllerImpl struct {
	service Service
}

func (c *ControllerImpl) GetGradesBySubject(ctx *fiber.Ctx) error {
	abbr := ctx.Query("abbr")
	number := ctx.Query("number")
	var distribution GradesDistribution
	err := c.service.FindGradesBySubject(abbr, number, &distribution)

	if err != nil {
		return err
	}

	return ctx.JSON(common.NewSuccessResponse(distribution))
}

func (c *ControllerImpl) GetCourseByAbbrAndNumber(ctx *fiber.Ctx) error {
	abbr := ctx.Query("abbr")
	number := ctx.Query("number")

	var course model.Course
	err := c.service.FindCourseBySubject(abbr, number, &course)
	if err != nil {
		return err
	}

	return ctx.JSON(common.NewSuccessResponse(NewOverviewFromCourse(course)))
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

	return ctx.JSON(common.NewSuccessResponse(course))
}
