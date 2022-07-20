package module2

import (
	"errors"
	"time"
)

// START Initial code

// Budget stores Budget information
type Budget struct {
	Max   float32
	Items []Item
}

// Item stores Item information
type Item struct {
	Description string
	Price       float32
}

var report map[time.Month]*Budget

// InitializeReport creates an empty map
// to store each budget
func InitializeReport() {
	report = make(map[time.Month]*Budget)
}

func init() {
	InitializeReport()
}

// CurrentCost returns how much we've added
// to the current budget
func (b Budget) CurrentCost() float32 {
	var sum float32
	for _, item := range b.Items {
		sum += item.Price
	}
	return sum
}

var errDoesNotFitBudget = errors.New("Item does not fit the budget")

var errReportIsFull = errors.New("Report is full")

var errDuplicateEntry = errors.New("Cannot add duplicate entry")

// END Initial code

// START Project code

// AddItem adds an item to the current budget
func (b *Budget) AddItem(description string, price float32) error {
	if b.CurrentCost()+price > b.Max {
		return errDoesNotFitBudget
	}

	newItem := Item{
		Description: description,
		Price:       price,
	}
	b.Items = append(b.Items, newItem)

	return nil
}

// RemoveItem removes a given item from the current budget
func (b *Budget) RemoveItem(description string) {
	for i := range b.Items {
		if b.Items[i].Description == description {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			break
		}
	}
}

// CreateBudget creates a new budget with a specified max
func CreateBudget(month time.Month, max float32) (*Budget, error) {
	var newBudget *Budget
	if len(report) >= 12 {
		return nil, errReportIsFull
	}
	if _, hasEntry := report[month]; hasEntry {
		return nil, errDuplicateEntry
	}
	newBudget = &Budget{
		Max: max,
	}
	report[month] = newBudget

	return newBudget, nil
}

// GetBudget returns budget for given month
func GetBudget(month time.Month) *Budget {
	if budget, ok := report[month]; ok {
		return budget
	}
	return nil
}

// END Project code
