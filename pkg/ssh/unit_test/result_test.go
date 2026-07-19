package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh/result"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestResult(t *testing.T) {
	assert.Any(
		t,
		&result.Result{
			OutputString: "Alfa",
			ErrorString:  "Bravo",
			Exit:         1,
			Error:        nil,
		},
		result.New(upper.Alfa, upper.Bravo, 1, nil),
	)
}
