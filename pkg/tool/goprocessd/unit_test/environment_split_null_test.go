package unit_test

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"strings"
	"testing"
)

func TestSplitNull(t *testing.T) {
	input := "KEY=value\x00OTHER=two\x00"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 2, len(tokens))
	assert.String(t, "KEY=value", tokens[0])
	assert.String(t, "OTHER=two", tokens[1])
}

func TestSplitNullEmbeddedNewline(t *testing.T) {
	input := "KEY=line1\nline2\x00OTHER=value\x00"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 2, len(tokens))
	assert.String(t, "KEY=line1\nline2", tokens[0])
	assert.String(t, "OTHER=value", tokens[1])
}

func TestSplitNullNoTrailingNull(t *testing.T) {
	input := "KEY=value"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 1, len(tokens))
	assert.String(t, "KEY=value", tokens[0])
}
