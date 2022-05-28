package organisation

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handle must start with @ in text")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "123"
}

func (p *Person) TwitterHandler() string {
	return p.twitterHandler
}
