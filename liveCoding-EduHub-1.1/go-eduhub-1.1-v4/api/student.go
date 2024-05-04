package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	GetStudents(c *gin.Context)
	GetStudentByID(c *gin.Context)
}

type studentAPI struct {
	studentRepo repo.StudentRepository
}

func NewStudentAPI(studentRepo repo.StudentRepository) *studentAPI {
	return &studentAPI{studentRepo}
}

func (s *studentAPI) AddStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	if err := s.studentRepo.Store(&student); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "add student success"})
}

func (s *studentAPI) GetStudents(c *gin.Context) {
	students, err := s.studentRepo.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (s *studentAPI) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	status, response := repository.GetStudentByID(studentID, s.studentRepo)
	c.JSON(status, response)
}
