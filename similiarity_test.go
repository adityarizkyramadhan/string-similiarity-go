package similiarity_test

import (
	"testing"

	similiarity "github.com/adityarizkyramadhan/string-similiarity-go"
)

func TestCompareTwoStrings(t *testing.T) {
	// Test cases
	tests := []struct {
		first, second string
	}{
		{"", ""},
		{"a", "a"},
		{"a", "b"},
		{"a", "ab"},
		{"ab", "ab"},
		{"ab", "ba"},
	}

	// Run tests
	for _, test := range tests {
		actual := similiarity.CompareTwoStrings(test.first, test.second)
		// assert agar tidak error gitu aja
		if actual < 0 || actual > 1 {
			t.Errorf("CompareTwoStrings(%q, %q) = %f; want a value between 0 and 1", test.first, test.second, actual)
		}
	}
}

func TestFindBestMatch(t *testing.T) {
	// Test cases
	tests := []struct {
		mainString    string
		targetStrings []string
	}{
		{"", []string{}},
		{"a", []string{"a"}},
		{"a", []string{"a", "b"}},
		{"ab", []string{"ab", "ba"}},
	}

	// Run tests
	for _, test := range tests {
		ratings, bestMatch := similiarity.FindBestMatch(test.mainString, test.targetStrings)
		if len(ratings) != len(test.targetStrings) {
			t.Errorf("FindBestMatch(%q, %q) = %v; want %v", test.mainString, test.targetStrings, ratings, test.targetStrings)
		}
		if bestMatch["target"] != test.targetStrings[0] {
			t.Errorf("FindBestMatch(%q, %q) = %v; want %v", test.mainString, test.targetStrings, bestMatch, test.targetStrings[0])
		}
	}
}
