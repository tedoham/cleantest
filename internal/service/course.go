package service

import (
	"github.com/tedoham/cleantest/internal/dto"
	"github.com/tedoham/cleantest/internal/repository"
)

type CourseService interface {
	GetCourse(courseID int64) (*dto.Course, error)
	// CreateCourse(course dto.Course) (*int64, error)
	// UpdateCourse(course dto.Course) (*dto.Course, error)
	// DeleteCourse(courseID int64, userID int64) error
}

type courseService struct {
	dao repository.DAO
}

func NewCourseService(dao repository.DAO) CourseService {
	return &courseService{dao: dao}
}

func (c *courseService) GetCourse(courseID int64) (*dto.Course, error) {
	course, err := c.dao.NewCourseQuery().GetCourse(courseID)
	if err != nil {
		return nil, err
	}

	fullCourse := dto.Course{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		Price:       course.Price,
		UserID:      course.UserID,
	}

	return &fullCourse, nil
}
