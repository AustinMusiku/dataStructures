package HashMap

import (
	"errors"
	"testing"
)

type Person struct {
	name string
	age  int
}

const (
	DefaultSize          = 23
	DefaultMaxLoadFactor = 0.75
)

func TestHashMap(t *testing.T) {
	t.Parallel()

	people := []Person{
		{"Rick", 40},
		{"Daryl", 30},
		{"Glenn", 25},
		{"Shane", 40},
		{"Tara", 25},
	}

	t.Run("Create hashmap", func(t *testing.T) {
		t.Parallel()

		hashmap := NewHashMap[int, Person](DefaultSize, DefaultMaxLoadFactor)

		if hashmap.size != DefaultSize {
			t.Errorf("Expected size to be %v, got %v", DefaultSize, hashmap.size)
		}
		if hashmap.maxLoadFactor != DefaultMaxLoadFactor {
			t.Errorf("Expected maxLoadFactor to be %v, got %v", DefaultMaxLoadFactor, hashmap.maxLoadFactor)
		}
		if hashmap.itemsCount != 0 {
			t.Errorf("Expected itemsCount to be 0, got %v", hashmap.itemsCount)
		}
	})

	t.Run("Add new item to hashmap", func(t *testing.T) {
		t.Parallel()

		hashmap := NewHashMap[string, Person](DefaultSize, DefaultMaxLoadFactor)

		hashmap.Set("Rick", people[0])
		rick, err := hashmap.Get("Rick")

		if hashmap.itemsCount != 1 {
			t.Errorf("Expected itemsCount to be 1, got %v", hashmap.itemsCount)
		}

		if err != nil {
			t.Errorf("Expected err to be nil, got %v", err)
		}

		if rick.key != "Rick" {
			t.Errorf("Expected key to be %v, got %v", "Rick", rick.key)
		}

		if rick.value != people[0] {
			t.Errorf("Expected value to be %v, got %v", people[0], rick.value)
		}
	})

	t.Run("Get item from hashmap", func(t *testing.T) {
		t.Parallel()

		type testGet struct {
			key           string
			expectedValue Person
			expectedErr   error
		}

		testCases := []testGet{
			{
				key:           "Rick",
				expectedValue: Person{"Rick", 40},
				expectedErr:   nil,
			},
			{
				key:           "Carol",
				expectedValue: Person{},
				expectedErr:   errors.New("key does not exist in map"),
			},
		}

		hashmap := NewHashMap[string, Person](DefaultSize, DefaultMaxLoadFactor)

		for _, person := range people {
			hashmap.Set(person.name, person)
		}

		for _, testCase := range testCases {
			value, err := hashmap.Get(testCase.key)

			if err != nil {
				if err.Error() != testCase.expectedErr.Error() {
					t.Errorf("Expected error to be %v, got %v", testCase.expectedErr, err)
				}
			} else {
				if value.value != testCase.expectedValue {
					t.Errorf("Expected value to be %v, got %v", testCase.expectedValue, value)
				}
			}
		}
	})

	t.Run("Delete item from hashmap", func(t *testing.T) {
		t.Parallel()

		type testDelete struct {
			key           string
			expectedValue Person
			expectedErr   error
		}

		testCases := []testDelete{
			{
				key:           "Rick",
				expectedValue: Person{"Rick", 40},
				expectedErr:   nil,
			},
			{
				key:           "Carol",
				expectedValue: Person{},
				expectedErr:   errors.New("key does not exist in map"),
			},
		}

		hashmap := NewHashMap[string, Person](DefaultSize, DefaultMaxLoadFactor)

		for _, person := range people {
			hashmap.Set(person.name, person)
		}

		for _, testCase := range testCases {
			value, err := hashmap.Delete(testCase.key)

			if err != nil {
				if err.Error() != testCase.expectedErr.Error() {
					t.Errorf("Expected error to be %v, got %v", testCase.expectedErr, err)
				}
			} else {
				if value.value != testCase.expectedValue {
					t.Errorf("Expected value to be %v, got %v", testCase.expectedValue, value)
				}

				if hashmap.itemsCount != 4 {
					t.Errorf("Expected itemsCount to be 4, got %v", hashmap.itemsCount)
				}
			}
		}
	})

	t.Run("Resize hashmap", func(t *testing.T) {
		t.Parallel()

		hashmapSize := 4
		hashmap := NewHashMap[string, Person](hashmapSize, DefaultMaxLoadFactor)

		hashmap.Resize(11)

		if hashmap.size != 11 {
			t.Errorf("Expected size to be 11, got %v", hashmap.size)
		}

		if hashmap.initSize != 4 {
			t.Errorf("Expected initSize to be 4, got %v", hashmap.initSize)
		}
	})

	t.Run("Reset hashmap", func(t *testing.T) {
		t.Parallel()

		hashmap := NewHashMap[string, Person](DefaultSize, DefaultMaxLoadFactor)

		for _, person := range people {
			hashmap.Set(person.name, person)
		}

		hashmap.Reset()

		if hashmap.size != DefaultSize {
			t.Errorf("Expected size to be %v, got %v", DefaultSize, hashmap.size)
		}

		if hashmap.itemsCount != 0 {
			t.Errorf("Expected itemsCount to be 0, got %v", hashmap.itemsCount)
		}
	})
}
