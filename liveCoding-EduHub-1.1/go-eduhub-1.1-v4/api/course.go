package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseAPI interface {
	AddCourse(c *gin.Context)
}

type courseAPI struct {
	courseRepo repository.CourseRepository
}

func NewCourseAPI(courseRepo repository.CourseRepository) *courseAPI {
	return &courseAPI{courseRepo}
}

func (cr *courseAPI) AddCourse(c *gin.Context) {
	var course model.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cr.courseRepo.Store(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "add course success"})
}
