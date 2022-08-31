package course

type Controller interface {
	GetCourseByName(name string, number int) (*Course, error)
}

type ControllerImpl struct {
	service Service
}
