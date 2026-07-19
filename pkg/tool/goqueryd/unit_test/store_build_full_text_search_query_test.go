package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func TestBuildFTSQuerySimple(t *testing.T) {
	result := store.BuildFullTextSearchQuery("hello world")
	assert.String(t, `"hello"* AND "world"*`, result)
}

func TestBuildFTSQueryNegation(t *testing.T) {
	result := store.BuildFullTextSearchQuery("hello -world")
	assert.String(t, `"hello"* NOT "world"*`, result)
}

func TestBuildFTSQueryQuotedPhrase(t *testing.T) {
	result := store.BuildFullTextSearchQuery(`"exact phrase"`)
	assert.String(t, `"exact phrase"`, result)
}

func TestBuildFTSQueryHyphenated(t *testing.T) {
	result := store.BuildFullTextSearchQuery("multi-agent")
	assert.String(t, `"multi agent"`, result)
}

func TestBuildFTSQueryNegatedHyphenated(t *testing.T) {
	result := store.BuildFullTextSearchQuery("-multi-agent search")
	assert.String(t, `"search"* NOT "multi agent"`, result)
}

func TestBuildFTSQueryEmpty(t *testing.T) {
	result := store.BuildFullTextSearchQuery("")
	assert.String(t, "", result)
}

func TestBuildFTSQueryOnlyNegation(t *testing.T) {
	result := store.BuildFullTextSearchQuery("-excluded")
	assert.String(t, "", result)
}

func TestBuildFTSQuerySpecialCharacters(t *testing.T) {
	result := store.BuildFullTextSearchQuery("hello! @world#")
	assert.String(t, `"hello"* AND "world"*`, result)
}
