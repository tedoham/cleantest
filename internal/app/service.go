package app

import "github.com/tedoham/cleantest/internal/service"

type MicroserviceServer struct {
	courseService service.CourseService
}

func NewMicroservice(
	courseService service.CourseService,
) *MicroserviceServer {
	return &MicroserviceServer{
		courseService: courseService,
	}
}
