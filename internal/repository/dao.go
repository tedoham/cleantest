package repository

type DAO interface {
	NewCourseQuery() CourseQuery
}

type dao struct{}
