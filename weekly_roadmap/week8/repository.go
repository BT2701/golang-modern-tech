package week8

import (
	"encoding/json"
	"modern-tech/weekly_roadmap/week3"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewStudentRepository(db *gorm.DB, rdb *redis.Client) *StudentRepository {
	return &StudentRepository{db: db, rdb: rdb}
}

func (r *StudentRepository) GetAllStudents() ([]week3.Student, error) {
	var students []week3.Student

	// Kiểm tra cache Redis
	cacheKey := "students"
	cachedData, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache không tồn tại, lấy từ DB
		if err := r.db.Find(&students).Error; err != nil {
			return nil, err
		}

		// Lưu vào Redis cache
		data, _ := json.Marshal(students)
		r.rdb.Set(ctx, cacheKey, data, time.Minute*10) // Cache 10 phút
	} else if err != nil {
		return nil, err
	} else {
		// Lấy từ cache
		json.Unmarshal([]byte(cachedData), &students)
	}

	return students, nil
}

func (r *StudentRepository) AddStudent(student *week3.Student) error {
	if err := r.db.Create(student).Error; err != nil {
		return err
	}

	// Xóa cache cũ
	r.rdb.Del(ctx, "students")
	return nil
}

func (r *StudentRepository) UpdateStudent(student *week3.Student) error {
	if err := r.db.Save(student).Error; err != nil {
		return err
	}

	// Xóa cache cũ
	r.rdb.Del(ctx, "students")
	return nil
}

func (r *StudentRepository) DeleteStudent(id int) error {
	if err := r.db.Delete(&week3.Student{}, id).Error; err != nil {
		return err
	}

	// Xóa cache cũ
	r.rdb.Del(ctx, "students")
	return nil
}

func (r *StudentRepository) IncrementAPICount(endpoint string) {
	r.rdb.Incr(ctx, "api_count:"+endpoint)
}

func (r *StudentRepository) GetAPIMetrics() (map[string]int, error) {
	keys, err := r.rdb.Keys(ctx, "api_count:*").Result()
	if err != nil {
		return nil, err
	}

	metrics := make(map[string]int)
	for _, key := range keys {
		count, _ := r.rdb.Get(ctx, key).Int()
		metrics[key] = count
	}

	return metrics, nil
}
