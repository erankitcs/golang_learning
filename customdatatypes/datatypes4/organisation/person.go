package organisation

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type europeanUnionIdentifier struct {
	id      string
	country []string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: []string{country},
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	First string
	Last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.First, n.Last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	First          string
	Last           string
	twitterHandler TwitterHandler
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			First: firstName,
			Last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handle must start with @ in text")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) ID() string {
	return "123"
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
