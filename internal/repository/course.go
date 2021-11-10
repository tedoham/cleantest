package repository

import "github.com/tedoham/cleantest/internal/entities"

type CourseQuery interface {
	GetCourse(id int64) (*entities.Course, error)
	CreateCourse(course entities.Course) (*int64, error)
	UpdateCourse(course entities.Course) (*entities.Course, error)
	DeleteCourse(id int64) error
}
