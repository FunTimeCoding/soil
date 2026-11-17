package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...service_tester.Option,
) *Tester {
	t.Helper()

	return &Tester{
		Tester: service_tester.New(
			t,
			append(
				[]service_tester.Option{
					service_tester.WithInstances(
						constant.FixtureInstance,
						constant.FixtureOtherInstance,
					),
				},
				o...,
			)...,
		),
		t: t,
	}
}
