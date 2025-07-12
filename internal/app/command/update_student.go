package command

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type UpdateStudentCommand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Tel   string `json:"tel"`
	Major string `json:"major"`
}

type UpdateStudentHandler struct {
	repo domain.Repository
}

func NewUpdateStudentHandler(repo domain.Repository) *UpdateStudentHandler {
	return &UpdateStudentHandler{repo: repo}
}

func (h *UpdateStudentHandler) Handle(ctx context.Context, cmd UpdateStudentCommand) (*domain.Student, error) {
	student, err := domain.NewStudent(cmd.ID, cmd.Name, cmd.Age, cmd.Tel, cmd.Major)
	if err != nil {
		return nil, err
	}
	return h.repo.Update(ctx, cmd.ID, student)
}
