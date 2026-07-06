//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestParseExpandedQueries(t *testing.T) {
	text := "lex: authentication login\nvec: user auth flow\nhyde: The system uses JWT tokens.\n"
	result := parseExpandedQueries(text, "auth")
	assert.Count(t, 3, result)
	assert.String(t, "lex", result[0].Type)
	assert.String(t, "authentication login", result[0].Query)
	assert.String(t, "vec", result[1].Type)
	assert.String(t, "hyde", result[2].Type)
}

func TestParseExpandedQueriesFiltersDuplicate(t *testing.T) {
	text := "lex: auth\nvec: authentication flow\n"
	result := parseExpandedQueries(text, "auth")
	assert.Count(t, 1, result)
	assert.String(t, "vec", result[0].Type)
}

func TestParseExpandedQueriesIgnoresInvalidTypes(t *testing.T) {
	text := "lex: valid\nfoo: invalid\nvec: also valid\n"
	result := parseExpandedQueries(text, "original")
	assert.Count(t, 2, result)
	assert.String(t, "lex", result[0].Type)
	assert.String(t, "vec", result[1].Type)
}

func TestParseExpandedQueriesEmptyOutput(t *testing.T) {
	result := parseExpandedQueries("", "query")
	assert.Count(t, 0, result)
}

func TestParseExpandedQueriesSkipsEmptyContent(t *testing.T) {
	text := "lex: \nvec: something\n"
	result := parseExpandedQueries(text, "original")
	assert.Count(t, 1, result)
	assert.String(t, "vec", result[0].Type)
}
