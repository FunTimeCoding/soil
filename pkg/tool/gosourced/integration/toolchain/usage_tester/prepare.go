package usage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) Prepare() *service_tester.Tester {
	o.t.Helper()
	gorm := o.t.TempDir()
	testutil.CopyTree(
		o.t,
		service_tester.ServiceTestdata("service-usage/gorm"),
		gorm,
	)
	system_tester.New(o.t).Write(
		gorm,
		"go.mod",
		"module gorm.io/gorm\n\ngo 1.26.3\n",
	)

	return service_tester.New(
		o.t,
		service_tester.WithFixture(
			"service-usage/src",
			"require gorm.io/gorm v0.0.0",
			fmt.Sprintf("replace gorm.io/gorm => %s", gorm),
		),
	)
}
