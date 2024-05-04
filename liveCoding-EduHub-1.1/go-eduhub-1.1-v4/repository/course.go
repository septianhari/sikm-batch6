package repository

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type CourseRepository interface {
	FetchByID(id int) (*model.Course, error)
	Store(course *model.Course) error
	ResetCourseRepo()
}

type courseRepository struct {
	courses []model.Course
}

func NewCourseRepo() *courseRepository {
	return &courseRepository{}
}

func (c *courseRepository) FetchByID(id int) (*model.Course, error) {
	for _, course := range c.courses {
		if course.ID == id {
			return &course, nil
		}
	}
	return nil, fmt.Errorf("course with id %d not found", id)
}

func (c *courseRepository) Store(course *model.Course) error {
	c.courses = append(c.courses, *course)
	return nil
}

func (c *courseRepository) ResetCourseRepo() {
	c.courses = []model.Course{}
}

func AddCourse(courseJSON []byte, courseRepo CourseRepository) (int, interface{}) {
	var course model.Course
	err := json.Unmarshal(courseJSON, &course)
	if err != nil {
		return http.StatusBadRequest, map[string]string{"error": err.Error()}
	}

	err = courseRepo.Store(&course)
	if err != nil {
		return http.StatusInternalServerError, map[string]string{"error": err.Error()}
	}

	response := map[string]string{"message": "add course success"}
	return http.StatusOK, response
}
