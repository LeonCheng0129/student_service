package server

import (
	"github.com/LeonCheng0129/student_service/internal/app/command"
	"github.com/LeonCheng0129/student_service/internal/app/query"
	"github.com/LeonCheng0129/student_service/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	createStudentHandler *command.CreateStudentHandler
	updateStudentHandler *command.UpdateStudentHandler
	getStudentHandler    *query.GetStudentHandler
	getAllStudentHandler *query.GetAllStudentsHandler
	deleteStudentHandler *command.DeleteStudentHandler
}

func NewServer(repo domain.Repository) *Server {
	return &Server{
		createStudentHandler: command.NewCreateStudentHandler(repo),
		updateStudentHandler: command.NewUpdateStudentHandler(repo),
		getStudentHandler:    query.NewGetStudentHandler(repo),
		getAllStudentHandler: query.NewGetAllStudentsHandler(repo),
		deleteStudentHandler: command.NewDeleteStudentHandler(repo),
	}
}

func toHTTPStudent(s *domain.Student) Student {
	return Student{
		Id:    &s.ID,
		Age:   &s.Age,
		Name:  &s.Name,
		Tel:   &s.Tel,
		Major: &s.Major,
	}
}

func (s *Server) GetStudents(c *gin.Context) {
	students, err := s.getAllStudentHandler.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to retrieve students"})
		return
	}

	// convert students to server.Student
	httpStudents := make([]Student, len(students))
	for _, student := range students {
		httpStudents = append(httpStudents, toHTTPStudent(student))
	}

	c.JSON(http.StatusOK, httpStudents)
}

func (s *Server) PostStudents(c *gin.Context) {
	var reqBody PostStudentsJSONRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request body"})
		return
	}

	cmd := command.CreateStudentCommand{
		Name:  reqBody.Name,
		Age:   reqBody.Age,
		Tel:   reqBody.Tel,
		Major: reqBody.Major,
	}
	student, err := s.createStudentHandler.Handle(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, toHTTPStudent(student))
}

func (s *Server) DeleteStudentsId(c *gin.Context, id int) {
	cmd := command.DeleteStudentCommand{ID: id}
	if err := s.deleteStudentHandler.Handle(c.Request.Context(), cmd); err != nil {
		if _, ok := err.(*domain.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to delete student"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) GetStudentsId(c *gin.Context, id int) {
	q := query.GetStudentQuery{ID: id}
	student, err := s.getStudentHandler.Handle(c.Request.Context(), q)
	if err != nil {
		if _, ok := err.(*domain.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to retrieve student"})
		return
	}
	c.JSON(http.StatusOK, toHTTPStudent(student))
}

func (s *Server) PutStudentsId(c *gin.Context, id int) {
	var reqBody PutStudentsIdJSONRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request body"})
		return
	}

	cmd := command.UpdateStudentCommand{
		ID:    id,
		Name:  reqBody.Name,
		Age:   reqBody.Age,
		Tel:   reqBody.Tel,
		Major: reqBody.Major,
	}
	student, err := s.updateStudentHandler.Handle(c.Request.Context(), cmd)
	if err != nil {
		if _, ok := err.(*domain.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to update student"})
		return
	}
	c.JSON(http.StatusOK, toHTTPStudent(student))
}
