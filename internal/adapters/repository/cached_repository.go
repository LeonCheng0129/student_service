package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LeonCheng0129/student_service/internal/domain"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"time"

	_ "github.com/LeonCheng0129/student_service/internal/common/configs"
)

type CachedRepository struct {
	redisClient *redis.Client
	next        domain.Repository
}

func (r *CachedRepository) Get(ctx context.Context, id int) (*domain.Student, error) {
	key := fmt.Sprintf("student:%d", id)
	jsonData, err := r.redisClient.Get(ctx, key).Result()

	if err == nil {
		// Cache hit, unmarshal jsonData to domain.Student
		var student domain.Student
		_ = json.Unmarshal([]byte(jsonData), &student)
		return &student, nil
	}

	if err != redis.Nil {
		// If the error is not a cache miss, return the error
		logrus.Errorf("Error retrieving student from redis: %v", err)
		return nil, err
	}

	// Cache miss (err == redis.Nil)
	// retrieve from the next repository
	student, err := r.next.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Store the student back to cache
	studentJSON, _ := json.Marshal(student)
	r.redisClient.Set(ctx, key, studentJSON, 10*time.Minute)

	return student, nil
}

func (r *CachedRepository) GetAll(ctx context.Context) ([]*domain.Student, error) {
	key := "students:all"
	jsonData, err := r.redisClient.Get(ctx, key).Result()

	if err == nil {
		// Cache hit, unmarshal jsonData to []*domain.Student
		var students []*domain.Student
		_ = json.Unmarshal([]byte(jsonData), &students)
		return students, nil
	}

	if err != redis.Nil {
		// If the error is not a cache miss, return the error
		logrus.Errorf("Error retrieving all students from redis: %v", err)
		return nil, err
	}

	// Cache miss (err == redis.Nil)
	students, err := r.next.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Store the students back to cache
	studentsJSON, _ := json.Marshal(students)
	r.redisClient.Set(ctx, key, studentsJSON, 10*time.Minute)

	return students, nil
}

/*******************************************************************
  Write operations, data consistency need to be handled carefully
  manipulate database first, if success, then delete cache
********************************************************************/

func (r *CachedRepository) Create(ctx context.Context, student *domain.Student) (*domain.Student, error) {
	createdStudent, err := r.next.Create(ctx, student)
	if err != nil {
		return nil, err
	}

	// Store the created student in cache
	key := fmt.Sprintf("student:%d", createdStudent.ID)
	studentJSON, _ := json.Marshal(createdStudent)
	err = r.redisClient.Set(ctx, key, studentJSON, 10*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error caching student ID %d: %v", createdStudent.ID, err)
		return nil, fmt.Errorf("failed to cache student ID %d: %w", createdStudent.ID, err)
	}

	// delete the all students cache
	keyAll := "students:all"
	err = r.redisClient.Del(ctx, keyAll).Err()
	if err != nil {
		logrus.Errorf("Error deleting all students cache: %v", err)
		return nil, fmt.Errorf("failed to delete all students cache: %w", err)
	}

	return createdStudent, nil
}

func (r *CachedRepository) Update(ctx context.Context, id int, student *domain.Student) (*domain.Student, error) {
	updatedStudent, err := r.next.Update(ctx, id, student)
	if err != nil {
		return nil, err
	}

	// delete the outdated cache
	key := fmt.Sprintf("student:%d", id)
	err = r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to delete cache for student ID %d: %w", id, err)
	}

	// delete the all students cache
	// TODO logic of updating all students cache needs to be refined
	keyAll := "students:all"
	err = r.redisClient.Del(ctx, keyAll).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to delete all students cache: %w", err)
	}
	return updatedStudent, nil
}

func (r *CachedRepository) Delete(ctx context.Context, id int) error {
	err := r.next.Delete(ctx, id)
	if err != nil {
		return err
	}

	// delete the cache for the deleted student
	key := fmt.Sprintf("student:%d", id)
	err = r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete cache for student ID %d: %w", id, err)
	}

	// delete the all students cache
	keyAll := "students:all"
	err = r.redisClient.Del(ctx, keyAll).Err()
	if err != nil {
		return fmt.Errorf("failed to delete all students cache: %w", err)
	}

	return nil
}

func NewCachedRepository(next domain.Repository) *CachedRepository {
	addr := fmt.Sprintf("%s:%s", viper.GetString("redis.ip"), viper.GetString("redis.port"))
	log.Printf("Connecting to Redis at %s", addr)
	redisClient := redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", viper.GetString("redis.ip"), viper.GetString("redis.port")),
		PoolSize:        viper.GetInt("redis.pool_size"),
		MaxActiveConns:  viper.GetInt("redis.max_conn"),
		ConnMaxLifetime: time.Duration(viper.GetInt("redis.conn_timeout")) * time.Millisecond,
		ReadTimeout:     time.Duration(viper.GetInt("redis.read_timeout")) * time.Millisecond,
		WriteTimeout:    time.Duration(viper.GetInt("redis.write_timeout")) * time.Millisecond,
	})

	// test connection
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)

	}
	log.Printf("Connected to Redis, received: %s", pong)

	return &CachedRepository{
		redisClient: redisClient,
		next:        next,
	}
}
