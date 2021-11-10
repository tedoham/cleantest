package service

import "github.com/tedoham/cleantest/internal/dto"

type CourseService interface {
	GetCourse(courseID int64) (*dto.Course, error)
	CreateCourse(course dto.Course) (*int64, error)
	UpdateCourse(course dto.Course) (*dto.Course, error)
	DeleteCourse(courseID int64, userID int64) error
}
