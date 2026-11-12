package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"slices"
	"testing"
)

func entryNames(
	t *testing.T,
	directory string,
) []string {
	t.Helper()
	entries, e := os.ReadDir(directory)
	assert.FatalOnError(t, e)
	var result []string

	for _, v := range entries {
		result = append(result, v.Name())
	}

	slices.Sort(result)

	return result
}
