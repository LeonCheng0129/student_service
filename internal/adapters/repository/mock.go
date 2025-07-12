package repository

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type MockRepository struct {
	repo []*domain.Student
}

func (m *MockRepository) Get(ctx context.Context, id int) (*domain.Student, error) {
	for _, student := range m.repo {
		if student.ID == id {
			return student, nil
		}
	}
	return nil, &domain.NotFoundError{ID: id}
}

func (m *MockRepository) GetAll(ctx context.Context) ([]*domain.Student, error) {
	return m.repo, nil
}

func (m *MockRepository) Create(ctx context.Context, student *domain.Student) (*domain.Student, error) {
	for _, existingStudent := range m.repo {
		if existingStudent.ID == student.ID {
			return nil, &domain.AlreadyExistsError{ID: student.ID}
		}
	}
	m.repo = append(m.repo, student)
	return student, nil
}

func (m *MockRepository) Update(ctx context.Context, id int, student *domain.Student) (*domain.Student, error) {
	for _, existingStudent := range m.repo {
		if existingStudent.ID == id {
			existingStudent.Name = student.Name
			existingStudent.Age = student.Age
			existingStudent.Tel = student.Tel
			existingStudent.Major = student.Major
			return existingStudent, nil
		}
	}
	return nil, &domain.NotFoundError{ID: id}
}

func (m *MockRepository) Delete(ctx context.Context, id int) error {
	for i, student := range m.repo {
		if student.ID == id {
			m.repo = append(m.repo[:i], m.repo[i+1:]...)
			return nil
		}
	}
	return &domain.NotFoundError{ID: id}
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		repo: []*domain.Student{
			{ID: 1, Name: "John Doe", Age: 20, Tel: "123-456-7890", Major: "Computer Science"},
			{ID: 2, Name: "Jane Smith", Age: 22, Tel: "987-654-3210", Major: "Mathematics"},
			{ID: 3, Name: "Alice Johnson", Age: 21, Tel: "555-555-5555", Major: "Physics"},
		},
	}
}
