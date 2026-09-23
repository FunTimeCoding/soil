package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/unit/strings_tester"
	"testing"
)

func assertCompare(
	t *testing.T,
	expected strings_tester.Compare,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := strings.Compare(past, now)
	assert.Strings(t, notNil(expected.Add), add)
	assert.Strings(t, notNil(expected.Remove), remove)
	assert.Strings(t, notNil(expected.Stay), stay)
}
