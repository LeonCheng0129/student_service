package repository

import (
	"context"
	"fmt"
	_ "github.com/LeonCheng0129/student_service/internal/common/configs"
	"github.com/LeonCheng0129/student_service/internal/domain"
	"github.com/spf13/viper"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// studentModel is the model for gorm
type studentModel struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Age       int
	Name      string `gorm:"type:varchar(100);not null"`
	Tel       string `gorm:"type:varchar(11);unique"`
	Major     string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (studentModel) TableName() string {
	return "student"
}

// toDomain converts studentModel to domain.Student
func (s *studentModel) toDomain() *domain.Student {
	return &domain.Student{
		ID:    s.ID,
		Name:  s.Name,
		Age:   s.Age,
		Tel:   s.Tel,
		Major: s.Major,
	}
}

// fromDomain converts domain.Student to studentModel
func fromStudent(from *domain.Student) *studentModel {
	return &studentModel{
		ID:    from.ID,
		Name:  from.Name,
		Age:   from.Age,
		Tel:   from.Tel,
		Major: from.Major,
	}
}

type MySQLRepository struct {
	db *gorm.DB
}

func (m *MySQLRepository) Get(ctx context.Context, id int) (*domain.Student, error) {
	var model studentModel
	result := m.db.WithContext(ctx).First(&model, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, &domain.NotFoundError{ID: id}
		}
		return nil, result.Error
	}
	return model.toDomain(), nil
}

func (m *MySQLRepository) GetAll(ctx context.Context) ([]*domain.Student, error) {
	var models []studentModel
	result := m.db.WithContext(ctx).Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	students := make([]*domain.Student, 0, len(models))
	for _, model := range models {
		students = append(students, model.toDomain())
	}

	return students, nil
}

func (m *MySQLRepository) Create(ctx context.Context, student *domain.Student) (*domain.Student, error) {
	model := fromStudent(student)

	result := m.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model.toDomain(), nil
}

func (m *MySQLRepository) Update(ctx context.Context, id int, student *domain.Student) (*domain.Student, error) {
	modelToUpdate := fromStudent(student)

	result := m.db.WithContext(ctx).Model(&studentModel{}).Where("id = ?", id).Updates(modelToUpdate)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &domain.NotFoundError{ID: id}
	}

	return student, nil
}

func (m *MySQLRepository) Delete(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Delete(&studentModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &domain.NotFoundError{ID: id}
	}

	return nil
}

func NewMySQLRepository() (*MySQLRepository, error) {
	// get dsn from config
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	dbname := viper.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	log.Printf("!!! CRITICAL DEBUG: Attempting to connect with DSN: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &MySQLRepository{db: db}, nil
}
