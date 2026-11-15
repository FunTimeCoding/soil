package service_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func Prepare(
	t *testing.T,
	name string,
	extraModLines ...string,
) string {
	t.Helper()

	return testutil.PrepareTestPackage(
		t,
		ServiceTestdata(name),
		extraModLines...)
}
