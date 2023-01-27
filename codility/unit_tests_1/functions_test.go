package unittests1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostUnique(t *testing.T) {
	tests := map[string]struct {
		a, b   []int
		exp    int
		expErr error
	}{
		"a empty":    {a: nil, b: []int{1}, exp: 0, expErr: errors.New("input is empty")},
		"b empty":    {a: []int{1}, b: nil, exp: 0, expErr: errors.New("input is empty")},
		"a too long": {a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9}, b: []int{1}, exp: 0, expErr: errors.New("input too long")},
		"b too long": {a: []int{1}, b: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9}, exp: 0, expErr: errors.New("input too long")},
		"success":    {a: []int{1, 2, 1}, b: []int{2, 3, 4, 5}, exp: 4, expErr: nil},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := MostUnique(tt.a, tt.b)

			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.exp, got)
		})
	}
}

func TestValidate(t *testing.T) {
	var validPerson = Person{
		Name:    "Roman",
		Surname: "Odesiuk",
		Age:     27,
		Hobbies: []string{"programming"},
	}

	tests := map[string]struct {
		people []Person
		exp    []error
	}{
		"wrong name and hobby": {
			people: []Person{{
				Name:    "Lina 0",
				Surname: "Todo",
				Hobbies: []string{"random", ""},
			}},
			exp: []error{
				errors.New("Person 0: Name cannot contain spaces, Hobbies must be valid (each of them must contain at least one letter)"),
			},
		},
		"wrong age, name, surname, hobby": {
			people: []Person{{
				Name:    "Anna",
				Surname: "Boleyn",
				Age:     516,
			}, {
				Name:    "putin",
				Surname: "D I C K H E A D",
				Age:     70,
				Hobbies: []string{"lie", "lie", "lie", "lie"},
			}},
			exp: []error{
				errors.New("Person 0: Age cannot be higher 120"),
				errors.New("Person 1: Name must start with capital letter, Surname cannot contain spaces, Hobbies can be empty, but if it contains values then the values must be unique (case insensitive comparison)"),
			},
		},
		"empty name,surname and age lower than 0": {
			people: []Person{validPerson, {
				Name:    "",
				Surname: "",
				Age:     -2,
			}},
			exp: []error{
				errors.New("Person 1: Name is empty, Surname is empty, Age is lower than 0"),
			},
		},
		"success": {
			people: []Person{validPerson},
			exp:    nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.exp, Validate(tt.people))
		})
	}
}
