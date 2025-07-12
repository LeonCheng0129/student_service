package command

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type CreateStudentCommand struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Tel   string `json:"tel"`
	Major string `json:"major"`
}

type CreateStudentHandler struct {
	repo domain.Repository
}

func NewCreateStudentHandler(repo domain.Repository) *CreateStudentHandler {
	return &CreateStudentHandler{repo: repo}
}

func (h *CreateStudentHandler) Handle(ctx context.Context, cmd CreateStudentCommand) (*domain.Student, error) {
	return h.repo.Create(ctx, &domain.Student{
		ID:    -1, // ID will be set by the repository
		Name:  cmd.Name,
		Age:   cmd.Age,
		Tel:   cmd.Tel,
		Major: cmd.Major,
	})
}
