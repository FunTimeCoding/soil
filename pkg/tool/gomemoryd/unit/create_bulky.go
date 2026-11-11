package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"strings"
	"testing"
)

func createBulky(
	t *testing.T,
	o *service_tester.Tester,
) {
	t.Helper()

	for i := range 80 {
		name := fmt.Sprintf("bulky %02d %s", i, strings.Repeat("filler ", 70))
		_, e := o.Service.CreateMemory(scopedOption(name, ""))
		assert.FatalOnError(t, e)
	}
}
