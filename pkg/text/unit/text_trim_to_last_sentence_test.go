package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/text"
	"testing"
)

func TestTrimToLastSentenceSingleSentencePeriod(t *testing.T) {
	assert.String(t, "Hello world.", text.TrimToLastSentence("Hello world."))
}

func TestTrimToLastSentenceSingleSentenceExclamation(t *testing.T) {
	assert.String(t, "Hello world!", text.TrimToLastSentence("Hello world!"))
}

func TestTrimToLastSentenceSingleSentenceQuestionMark(t *testing.T) {
	assert.String(t, "Hello world?", text.TrimToLastSentence("Hello world?"))
}

func TestTrimToLastSentenceTwoSentences(t *testing.T) {
	assert.String(
		t,
		"First sentence.",
		text.TrimToLastSentence("First sentence. Incomplete second"),
	)
}

func TestTrimToLastSentenceMixedPunctuation(t *testing.T) {
	assert.String(
		t,
		"First! Second?",
		text.TrimToLastSentence("First! Second? Third incomplete"),
	)
}

func TestTrimToLastSentenceNoTerminator(t *testing.T) {
	assert.String(
		t,
		"This is incomplete",
		text.TrimToLastSentence("This is incomplete"),
	)
}

func TestTrimToLastSentenceEmpty(t *testing.T) {
	assert.String(t, "", text.TrimToLastSentence(""))
}

func TestTrimToLastSentenceSingleCharacter(t *testing.T) {
	assert.String(t, "a", text.TrimToLastSentence("a"))
}

func TestTrimToLastSentenceEndsWithPunctuation(t *testing.T) {
	assert.String(
		t,
		"Complete sentence.",
		text.TrimToLastSentence("Complete sentence."),
	)
}

func TestTrimToLastSentencePunctuationAtStart(t *testing.T) {
	assert.String(t, ".", text.TrimToLastSentence(".Start with period"))
}

func TestTrimToLastSentenceOnlyPunctuation(t *testing.T) {
	assert.String(t, ".", text.TrimToLastSentence(constant.Dot))
}

func TestTrimToLastSentenceNoFollowingText(t *testing.T) {
	assert.String(t, "Sentence.", text.TrimToLastSentence("Sentence."))
}

func TestTrimToLastSentenceMultiplePunctuationMarks(t *testing.T) {
	assert.String(
		t,
		"What?! Really.",
		text.TrimToLastSentence("What?! Really."),
	)
}

func TestTrimToLastSentenceAbbreviation(t *testing.T) {
	assert.String(t, "Dr.", text.TrimToLastSentence("Dr. Smith said hello"))
}

func TestTrimToLastSentenceDecimalNumber(t *testing.T) {
	assert.String(
		t,
		"The value is 3.",
		text.TrimToLastSentence("The value is 3.14 approximately"),
	)
}

func TestTrimToLastSentenceDomainWithPeriods(t *testing.T) {
	assert.String(
		t,
		"Visit example.",
		text.TrimToLastSentence("Visit example.com for more info"),
	)
}

func TestTrimToLastSentenceEllipsis(t *testing.T) {
	assert.String(t, "Well...", text.TrimToLastSentence("Well..."))
}

func TestTrimToLastSentenceTrailingWhitespace(t *testing.T) {
	assert.String(t, "Sentence.", text.TrimToLastSentence("Sentence.   "))
}

func TestTrimToLastSentenceNewline(t *testing.T) {
	assert.String(
		t,
		"First sentence.",
		text.TrimToLastSentence("First sentence.\nSecond incomplete"),
	)
}

func TestTrimToLastSentenceUnicode(t *testing.T) {
	assert.String(
		t,
		"Hello 世界.",
		text.TrimToLastSentence("Hello 世界. More text"),
	)
}
