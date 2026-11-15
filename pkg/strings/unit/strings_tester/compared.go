package strings_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func Compared(
	t *testing.T,
	expected Compare,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := strings.Compare(past, now)
	assert.Strings(t, notNil(expected.Add), add)
	assert.Strings(t, notNil(expected.Remove), remove)
	assert.Strings(t, notNil(expected.Stay), stay)
}
