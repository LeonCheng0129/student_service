// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Student defines model for Student.
type Student struct {
	// Age Age of the student
	Age *int `json:"age,omitempty"`

	// Id Unique identifier for the student
	Id *int `json:"id,omitempty"`

	// Major Major field of study of the student
	Major *string `json:"major,omitempty"`

	// Name Name of the student
	Name *string `json:"name,omitempty"`

	// Tel Telephone number of the student
	Tel *string `json:"tel,omitempty"`
}

// StudentInput defines model for StudentInput.
type StudentInput struct {
	// Age Age of the student
	Age int `json:"age"`

	// Major Major field of study of the student
	Major string `json:"major"`

	// Name Name of the student
	Name string `json:"name"`

	// Tel Telephone number of the student
	Tel string `json:"tel"`
}

// PostStudentsJSONRequestBody defines body for PostStudents for application/json ContentType.
type PostStudentsJSONRequestBody = StudentInput

// PutStudentsIdJSONRequestBody defines body for PutStudentsId for application/json ContentType.
type PutStudentsIdJSONRequestBody = StudentInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve all students information
	// (GET /students)
	GetStudents(c *gin.Context)
	// Add a new student
	// (POST /students)
	PostStudents(c *gin.Context)
	// Delete a specific student's information based on ID
	// (DELETE /students/{id})
	DeleteStudentsId(c *gin.Context, id int)
	// Retrieve a specific student's information based on ID
	// (GET /students/{id})
	GetStudentsId(c *gin.Context, id int)
	// Update a specific student's information based on ID
	// (PUT /students/{id})
	PutStudentsId(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetStudents operation middleware
func (siw *ServerInterfaceWrapper) GetStudents(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetStudents(c)
}

// PostStudents operation middleware
func (siw *ServerInterfaceWrapper) PostStudents(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostStudents(c)
}

// DeleteStudentsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteStudentsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteStudentsId(c, id)
}

// GetStudentsId operation middleware
func (siw *ServerInterfaceWrapper) GetStudentsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetStudentsId(c, id)
}

// PutStudentsId operation middleware
func (siw *ServerInterfaceWrapper) PutStudentsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutStudentsId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/students", wrapper.GetStudents)
	router.POST(options.BaseURL+"/students", wrapper.PostStudents)
	router.DELETE(options.BaseURL+"/students/:id", wrapper.DeleteStudentsId)
	router.GET(options.BaseURL+"/students/:id", wrapper.GetStudentsId)
	router.PUT(options.BaseURL+"/students/:id", wrapper.PutStudentsId)
}
