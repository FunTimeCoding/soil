package unit_test

import (
	"github.com/funtimecoding/soil/pkg/text"
	"testing"
)

func TestTrimToLastSentence(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{
			name:     "single_complete_sentence_with_period",
			input:    "Hello world.",
			expected: "Hello world.",
		},
		{
			name:     "single_complete_sentence_with_exclamation",
			input:    "Hello world!",
			expected: "Hello world!",
		},
		{
			name:     "single_complete_sentence_with_question_mark",
			input:    "Hello world?",
			expected: "Hello world?",
		},
		// Multiple sentences
		{
			name:     "two_sentences_trim_after_first",
			input:    "First sentence. Incomplete second",
			expected: "First sentence.",
		},
		{
			name:     "multiple_sentences_with_mixed_punctuation",
			input:    "First! Second? Third incomplete",
			expected: "First! Second?",
		},
		// Edge cases with no complete sentences
		{
			name:     "no_sentence_terminator",
			input:    "This is incomplete",
			expected: "This is incomplete",
		},
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "single_character",
			input:    "a",
			expected: "a",
		},
		// Punctuation at boundaries
		{
			name:     "ends_with_punctuation",
			input:    "Complete sentence.",
			expected: "Complete sentence.",
		},
		{
			name:     "punctuation_at_start",
			input:    ".Start with period",
			expected: ".",
		},
		{
			name:     "only_punctuation",
			input:    ".",
			expected: ".",
		},
		{
			name:     "punctuation_with_no_following_text",
			input:    "Sentence.",
			expected: "Sentence.",
		},
		// Special cases
		{
			name:     "multiple_punctuation_marks",
			input:    "What?! Really.",
			expected: "What?! Really.",
		},
		{
			name:     "abbreviation_with_period",
			input:    "Dr. Smith said hello",
			expected: "Dr.",
		},
		{
			name:     "decimal_number",
			input:    "The value is 3.14 approximately",
			expected: "The value is 3.",
		},
		{
			name:     "URL_with_periods",
			input:    "Visit example.com for more info",
			expected: "Visit example.",
		},
		{
			name:     "ellipsis_incomplete",
			input:    "Well...",
			expected: "Well...",
		},
		// Whitespace handling
		{
			name:     "trailing_whitespace_after_punctuation",
			input:    "Sentence.   ",
			expected: "Sentence.",
		},
		{
			name:     "sentence_with_newline",
			input:    "First sentence.\nSecond incomplete",
			expected: "First sentence.",
		},
		// Unicode/international
		{
			name:     "unicode_text",
			input:    "Hello 世界. More text",
			expected: "Hello 世界.",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				actual := text.TrimToLastSentence(tt.input)

				if actual != tt.expected {
					t.Errorf(
						"text.TrimToLastSentence(%q) = %q, want %q",
						tt.input,
						actual,
						tt.expected,
					)
				}
			},
		)
	}
}
