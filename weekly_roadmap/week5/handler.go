package week5

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *StudentService
}

func NewStudentHandler(service *StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// Lấy danh sách sinh viên
func (h *StudentHandler) GetAll(c *gin.Context) {
	students, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// Lấy thông tin sinh viên theo ID
func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

// Thêm sinh viên mới
func (h *StudentHandler) Add(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdStudent, err := h.service.Add(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdStudent)
}

// Cập nhật thông tin sinh viên
func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	student.ID = id
	updatedStudent, err := h.service.Update(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

// Xóa sinh viên
func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Đăng ký các route
func RegisterRoutes(router *gin.Engine, handler *StudentHandler) {
	students := router.Group("/students")
	{
		students.GET("", handler.GetAll)        // GET /students
		students.GET("/:id", handler.GetByID)   // GET /students/:id
		students.POST("", handler.Add)          // POST /students
		students.PUT("/:id", handler.Update)    // PUT /students/:id
		students.DELETE("/:id", handler.Delete) // DELETE /students/:id
	}
}
