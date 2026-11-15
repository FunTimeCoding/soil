package service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/save_tester"
	"strings"
)

func (o *Tester) CreateBulky() {
	o.t.Helper()

	for i := range 80 {
		name := fmt.Sprintf("bulky %02d %s", i, strings.Repeat("filler ", 70))
		_, e := o.Service.CreateMemory(save_tester.Scoped(name, ""))
		assert.FatalOnError(o.t, e)
	}
}
