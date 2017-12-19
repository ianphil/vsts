package tests

import (
	"fmt"
	"testing"
)

type Person struct {
	Firstname string
	Lastname  string
}

func (p Person) ToString() string {
	return fmt.Sprintf("%s, %s", p.Lastname, p.Firstname)
}

func TestToString(t *testing.T) {
	person := Person{"Ian", "Philpot"}
	result := person.ToString()

	if result != "Philpot, Ian" {
		t.Errorf("To string doesn't work")
	}
}
