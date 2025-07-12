package query

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type GetAllStudentsHandler struct {
	repo domain.Repository
}

func NewGetAllStudentsHandler(repo domain.Repository) *GetAllStudentsHandler {
	return &GetAllStudentsHandler{repo: repo}
}

func (h *GetAllStudentsHandler) Handle(ctx context.Context) ([]*domain.Student, error) {
	return h.repo.GetAll(ctx)
}
