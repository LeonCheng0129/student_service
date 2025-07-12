package command

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type DeleteStudentCommand struct {
	ID int
}

type DeleteStudentHandler struct {
	repo domain.Repository
}

func NewDeleteStudentHandler(repo domain.Repository) *DeleteStudentHandler {
	return &DeleteStudentHandler{repo: repo}
}

func (h *DeleteStudentHandler) Handle(ctx context.Context, cmd DeleteStudentCommand) error {
	return h.repo.Delete(ctx, cmd.ID)
}
