package Trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	t.Parallel()

	t.Run("Create trie", func(t *testing.T) {
		trie := NewTrie()

		if trie == nil {
			t.Error("Trie is nil")
		}
	})

	t.Run("Insert word", func(t *testing.T) {
		trie := NewTrie()

		inserted := trie.Insert("hey")
		found := trie.Search("hey")

		if !inserted {
			t.Error("Expected \"hey\" to be inserted")
		}

		if !found {
			t.Error("Expected \"hey\" to be found")
		}
	})

	t.Run("Search for word", func(t *testing.T) {
		type testSearch struct {
			insert   string
			search   string
			expected bool
		}

		testCases := []testSearch{
			{"", "", false},
			{"", "hey", false},
			{"hey", "hey", true},
			{"hey", "he", false},
			{"hey", "heyy", false},
			{"hey", "jey", false},
		}

		for _, testCase := range testCases {
			trie := NewTrie()

			trie.Insert(testCase.insert)

			ok := trie.Search(testCase.search)

			if ok != testCase.expected {
				t.Errorf("Expected %v, got %v", testCase.expected, ok)
			}
		}
	})

	t.Run("Delete word", func(t *testing.T) {
		type testDelete struct {
			insert   string
			delete   string
			expected bool
		}

		testCases := []testDelete{
			{"", "", false},
			{"", "hey", false},
			{"hey", "", false},
			{"hey", "hey", true},
		}

		for _, testCase := range testCases {
			trie := NewTrie()

			trie.Insert(testCase.insert)

			delete := trie.Delete(testCase.delete)

			if delete != testCase.expected {
				t.Errorf("Expected %v, got %v", testCase.expected, delete)
			}

			search := trie.Search(testCase.delete)

			if delete && search {
				t.Errorf("Expected %v to not be found in trie", testCase.delete)
			}
		}
	})
}
