package test

import (
	"testing"
	"wordcounter/utils"
	"wordcounter/parser"
)

func TestExtractWords(t *testing.T) {
	content := utils.Content{
		Title:       "Sample Title",
		Heading:     "Sample Heading",
		Description: "This is a sample description with valid words.",
	}

	expected := []string{"Sample", "Title", "Sample", "Heading", "This", "is", "a", "sample", "description", "with", "valid", "words"}
	got := parser.ExtractWords(content)

	if len(got) != len(expected) {
		t.Fatalf("expected %d words, got %d", len(expected), len(got))
	}

	for i, word := range expected {
		if got[i] != word {
			t.Errorf("expected %s, got %s", word, got[i])
		}
	}
}
