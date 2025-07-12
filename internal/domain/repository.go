package domain

import (
	"context"
)

// Repository is the interface that defines the methods for student repository.
type Repository interface {
	Get(ctx context.Context, id int) (*Student, error)
	GetAll(ctx context.Context) ([]*Student, error)
	Create(ctx context.Context, student *Student) (*Student, error)
	Update(ctx context.Context, id int, student *Student) (*Student, error)
	Delete(ctx context.Context, id int) error
}

// define some error types for repository operations
type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return "student not found with ID: " + string(e.ID)
}

type AlreadyExistsError struct {
	ID int
}

func (e *AlreadyExistsError) Error() string {
	return "student already exists with ID: " + string(e.ID)
}
