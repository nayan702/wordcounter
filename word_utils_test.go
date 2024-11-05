package test

import (
	"testing"
	"wordcounter/utils"
)

func TestCountValidWords(t *testing.T) {
	wordBank := map[string]struct{}{
		"valid":   {},
		"test":    {},
		"example": {},
	}

	tests := []struct {
		words    []string
		expected map[string]int
	}{
		{[]string{"valid", "VALID", "test", "test"}, map[string]int{"valid": 2, "test": 2}},
		{[]string{"invalid", "word"}, map[string]int{}},
	}

	for _, test := range tests {
		got := utils.CountValidWords(test.words, wordBank)
		if len(got) != len(test.expected) {
			t.Errorf("expected %d words, got %d", len(test.expected), len(got))
		}
		for word, count := range test.expected {
			if got[word] != count {
				t.Errorf("for word %s: expected %d, got %d", word, count, got[word])
			}
		}
	}
}
