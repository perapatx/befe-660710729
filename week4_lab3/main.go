package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:year`
	GPA   float64 `json:"gpa"`
}

// s is instance has all attribute of Student
func (s *Student) IsHornor() bool {
	return s.GPA >= 3.50
}
func (s *Student) Validate() error {

	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("Year Must between 1 and 4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("Gpa must between 0-4")
	}
	return nil
}

func main() {
	// var st Student = Student{ID: "1", Name: "Perapat", Email: "perapat@gmail.com", Year: 3, GPA: 3.8}
	students := []Student{
		{ID: "1", Name: "Perapat", Email: "perapat@gmail.com", Year: 3, GPA: 3.8},
		{ID: "2", Name: "Alice", Email: "perapat@gmail.com", Year: 3, GPA: 1.5},
		{ID: "3", Name: "Ruburt", Email: "perapat@gmail.com", Year: 5, GPA: 3},
		{ID: "4", Name: "Victoria", Email: "", Year: 3, GPA: 2.5},
		{ID: "5", Name: "Hambert", Email: "perapat@gmail.com", Year: 3, GPA: 8},
		{ID: "6", Name: "June", Email: "perapat@gmail.com", Year: 3, GPA: 3.4},
		{ID: "7", Name: "Jack", Email: "perapat@gmail.com", Year: 3, GPA: 1.0},
		{ID: "8", Name: "Mew", Email: "perapat@gmail.com", Year: 3, GPA: 2.8},
	}
	newStudent := Student{ID: "9", Name: "Judy", Email: "judy@gmail.com", Year: 3, GPA: 3.95}
	students = append(students, newStudent)

	for _, student := range students {
		fmt.Printf("Hornor: %v\n", student.IsHornor())
		fmt.Printf("Validation: %v\n", student.Validate())
	}

}
