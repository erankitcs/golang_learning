package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var gradeSum float32
	for _, grade := range s.Grades {
		gradeSum += grade.Score
	}
	return gradeSum / float32(len(s.Grades))

}

type Students []Student

func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("Student with ID: %v not found", id)

}

var (
	students      Students
	studentsMutex sync.Mutex
)

type GradeType string

const (
	GradeTest     = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz     = GradeType("Quiz")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
