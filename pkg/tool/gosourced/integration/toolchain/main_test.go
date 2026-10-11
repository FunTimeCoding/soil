package toolchain

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	errors.PanicOnError(os.RemoveAll(service_tester.IndexDirectory()))
	os.Exit(code)
}
