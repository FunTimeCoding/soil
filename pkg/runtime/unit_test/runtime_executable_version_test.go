package unit_test

import (
	"github.com/funtimecoding/soil/pkg/runtime"
	"testing"
)

func TestExecutableVersion(t *testing.T) {
	v := runtime.ExecutableVersion()

	if v == nil {
		t.Error("Expect version to be not nil")

		return
	}

	if v.Major < 1 {
		t.Errorf(
			"Expect major to be at least 1, got %d",
			v.Major,
		)
	}

	if v.Minor < 23 {
		t.Errorf(
			"Expect minor to be at least 23, got %d",
			v.Minor,
		)
	}
}
