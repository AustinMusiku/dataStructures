package BloomFilter

import "testing"

func TestBloomFilter(t *testing.T) {
	t.Parallel()

	testCases := []string{"rick grimmes", "carl grimmes", "daryl dixon", "michonne", "rosita espinosa", "glenn rhee", "maggie greene", "abraham ford", "eugene porter", "sasha williams", "tyreese williams", "beth greene", "hershel greene", "the governor", "merle dixon", "shane walsh", "dale horvath", "andrea", "lori grimmes", "judith grimmes", "morgan jones", "noah", "gabriel stokes", "spencer monroe", "aaron", "enid", "gareth", "deanna monroe", "douglas monroe", "jessie anderson", "ron anderson", "pete anderson", "nicholas", "denise cloyd", "olivia", "eric raleigh", "heath", "scott", "bob stookey", "tara chambler", "siddiq", "dwight", "simon", "gregory", "jerry", "ezequiel", "daniel", "lydia", "alpha", "beta", "ezekiel", "dante", "connie", "kelly", "luke", "yumiko", "magna", "shaw", "sebastian milton"}

	bf := NewBloomFilter[string](10000000)

	// Test for positives
	for _, testCase := range testCases {
		bf.Add(testCase)

		if !bf.Check(testCase) {
			t.Errorf("Expected \"%s\" to be in the filter, got false", testCase)
		}
	}

	// Test for negatives
	for _, testCase := range testCases {
		randomString := testCase + "testcase"
		if bf.Check(randomString) {
			t.Errorf("Expected \"%s\" to not be in the filter, got true", randomString)
		}
	}

	// clear
	bf.Reset()
	for _, testCase := range testCases {
		if bf.Check(testCase) {
			t.Errorf("Expected filter to be empty, got \"%s\"", testCase)
		}
	}
}
