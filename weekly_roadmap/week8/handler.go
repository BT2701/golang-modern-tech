package week8

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	repo *StudentRepository
}

func NewStudentHandler(repo *StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	h.repo.IncrementAPICount("GET /students")

	students, err := h.repo.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetMetrics(c *gin.Context) {
	metrics, err := h.repo.GetAPIMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}
