package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// Student struct
type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

func getStudents(c *gin.Context) {
	yearQuery := c.Query("Year")
	if yearQuery != "" {
		filter := []Student{}
		for _, student := range students {
			if fmt.Sprint(student.Year) == yearQuery {
				filter = append(filter, student)
			}
		}

		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)
}

func getStudent(input *gin.Context) {
	ID := input.Param("id")
	for _, student := range students {
		if student.ID == ID {
			input.JSON(http.StatusOK, student)
			return
		}
	}
	input.JSON(http.StatusNotFound, gin.H{"error": "please input student id"})
}

func createStudent(input *gin.Context) {
	var newStudent Student
	if err := input.ShouldBindJSON(&newStudent); err != nil { //ERROR
		input.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newStudent.Name == "" {
		input.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if newStudent.Year < 1 || newStudent.Year > 4 {
		input.JSON(http.StatusBadRequest, gin.H{"error": "year not correct"})
		return
	}
	//SUCCES
	newStudent.ID = fmt.Sprintf("%d", len(students)+1)
	students = append(students, newStudent)
	input.JSON(http.StatusOK, students)

}

func updateStudent(input *gin.Context) {
	id := input.Param("id")
	var UpdateStudent Student

	if err := input.ShouldBindJSON(&UpdateStudent); err != nil { //ERROR
		input.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == id {
			UpdateStudent.ID = id
			students[i] = UpdateStudent
			input.JSON(http.StatusOK, UpdateStudent)
			return
		}
	}
	input.JSON(http.StatusNotFound, gin.H{"Error": "Not found"})
}

func deleteStudent(input *gin.Context) {
	id := input.Param("id")
	for i, student := range students {
		if student.ID == id {
			students = slices.Delete(students, i, i+1)
			input.JSON(http.StatusOK, gin.H{"Success": "Successssss"})
			return
		}
	}
	input.JSON(http.StatusNotFound, gin.H{"Error": "Not found"})
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"massage": "healthy"})
		fmt.Println()
	})
	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents)
		api.GET("/students/:id", getStudent)
		api.POST("/students", createStudent)
		api.PUT("/students/:id", updateStudent)
		api.DELETE("/students/:id", deleteStudent)
		fmt.Println()
	}

	r.Run(":8080")
}
