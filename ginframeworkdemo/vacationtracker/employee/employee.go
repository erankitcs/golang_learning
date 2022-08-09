package employee

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimeOff   []TimeOff
}

type TimeOff struct {
	Type      TimeoffType
	Amount    float32   `form:"amount" json:"amount" binding:"required"`
	StartDate time.Time `form:"date" json:"date" binding:"required" time_format:"2006-01-02"`
	Status    TimeoffStatus
}

type TimeoffStatus string
type TimeoffType string

const (
	TimeoffStatusRequested TimeoffStatus = "Requested"
	TimeoffStatusScheduled TimeoffStatus = "Scheduled"
	TimeoffStatusTaken     TimeoffStatus = "Taken"
)

const (
	TimeoffTypeHoliday TimeoffType = "Holiday"
	TimeoffTypePTO     TimeoffType = "PTO"
)

func Get(id int) (*Employee, error) {
	for i := range employees {
		if employees[i].ID == id {
			return &employees[i], nil
		}
	}
	return nil, fmt.Errorf("employee with id %v not found", id)
}

func GetAll() []Employee {
	return employees
}
