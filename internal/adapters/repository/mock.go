package repository

import (
	"context"
	"github.com/LeonCheng0129/student_service/internal/domain"
)

type MockRepository struct {
	repo   map[int]*domain.Student
	nextID int
}

func (m *MockRepository) Get(_ context.Context, id int) (*domain.Student, error) {
	student, exits := m.repo[id]
	if !exits {
		return nil, &domain.NotFoundError{ID: id}
	}
	return student, nil
}

func (m *MockRepository) GetAll(_ context.Context) ([]*domain.Student, error) {
	students := make([]*domain.Student, 0, len(m.repo))
	for _, student := range m.repo {
		students = append(students, student)
	}
	return students, nil
}

func (m *MockRepository) Create(_ context.Context, student *domain.Student) (*domain.Student, error) {
	student.ID = m.nextID
	if _, exists := m.repo[student.ID]; exists {
		return nil, &domain.AlreadyExistsError{ID: student.ID}
	}
	m.repo[student.ID] = student
	m.nextID++
	return student, nil
}

func (m *MockRepository) Update(_ context.Context, id int, student *domain.Student) (*domain.Student, error) {
	if _, exists := m.repo[id]; !exists {
		return nil, &domain.NotFoundError{ID: id}
	}
	// Update the student details
	m.repo[id] = student
	return student, nil
}

func (m *MockRepository) Delete(_ context.Context, id int) error {
	if _, exists := m.repo[id]; !exists {
		return &domain.NotFoundError{ID: id}
	}
	delete(m.repo, id)
	return nil
}

func NewMockRepository() *MockRepository {
	repo := map[int]*domain.Student{
		1: {ID: 1, Name: "Alice", Age: 20, Tel: "1234567890", Major: "Computer Science"},
		2: {ID: 2, Name: "Bob", Age: 22, Tel: "0987654321", Major: "Mathematics"},
		3: {ID: 3, Name: "Charlie", Age: 21, Tel: "1122334455", Major: "Physics"},
		4: {ID: 4, Name: "David", Age: 23, Tel: "5566778899", Major: "Chemistry"},
		5: {ID: 5, Name: "Eve", Age: 19, Tel: "6677889900", Major: "Biology"},
	}
	return &MockRepository{
		repo:   repo,
		nextID: len(repo) + 1,
	}
}
