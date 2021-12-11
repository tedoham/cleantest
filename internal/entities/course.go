package entities

const CourseTableName = "course"

type Course struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
}
