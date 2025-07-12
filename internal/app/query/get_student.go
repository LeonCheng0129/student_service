package query

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type GetStudentQuery struct {
	ID int
}

type GetStudentHandler struct {
	repo domain.Repository
}

func NewGetStudentHandler(repo domain.Repository) *GetStudentHandler {
	return &GetStudentHandler{repo: repo}
}

func (h *GetStudentHandler) Handle(ctx context.Context, query GetStudentQuery) (*domain.Student, error) {
	student, err := h.repo.Get(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return student, nil
}
