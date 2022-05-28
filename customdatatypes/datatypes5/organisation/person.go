package organisation

import (
	"errors"
	"fmt"
	"strconv"
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

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: []string{country},
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: []string{country},
		}
	case europeanUnionIdentifier:
		return v
	case Person:
		euId, _ := v.Citizen.(europeanUnionIdentifier)
		return euId
	default:
		panic("Using invalid type initialize EU identifier")
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

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
