package architecture

import "fmt"

// Person is how the architecture package stores a person
type Person struct {
	First string
}

// Accessor is how to store or retreive a person
// When retriving a person, if they do not exist, return the zero value
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

type PersonService struct {
	a Accessor
}

// Create a new PersonService to allow the ability to accessthe data
func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No Person with n of %d", n)
	}
	return p, nil
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

