//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestBuildFTSQuerySimple(t *testing.T) {
	result := buildFullTextSearchQuery("hello world")
	assert.String(t, `"hello"* AND "world"*`, result)
}

func TestBuildFTSQueryNegation(t *testing.T) {
	result := buildFullTextSearchQuery("hello -world")
	assert.String(t, `"hello"* NOT "world"*`, result)
}

func TestBuildFTSQueryQuotedPhrase(t *testing.T) {
	result := buildFullTextSearchQuery(`"exact phrase"`)
	assert.String(t, `"exact phrase"`, result)
}

func TestBuildFTSQueryHyphenated(t *testing.T) {
	result := buildFullTextSearchQuery("multi-agent")
	assert.String(t, `"multi agent"`, result)
}

func TestBuildFTSQueryNegatedHyphenated(t *testing.T) {
	result := buildFullTextSearchQuery("-multi-agent search")
	assert.String(t, `"search"* NOT "multi agent"`, result)
}

func TestBuildFTSQueryEmpty(t *testing.T) {
	result := buildFullTextSearchQuery("")
	assert.String(t, "", result)
}

func TestBuildFTSQueryOnlyNegation(t *testing.T) {
	result := buildFullTextSearchQuery("-excluded")
	assert.String(t, "", result)
}

func TestBuildFTSQuerySpecialCharacters(t *testing.T) {
	result := buildFullTextSearchQuery("hello! @world#")
	assert.String(t, `"hello"* AND "world"*`, result)
}
