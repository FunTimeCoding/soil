package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Integer[T integerKind](
	t *testing.T,
	expect T,
	actual T,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
