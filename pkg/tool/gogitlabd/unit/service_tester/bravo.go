package service_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func (o *Tester) Bravo() *mock_client.Client {
	o.t.Helper()

	return o.Instance(constant.FixtureOtherInstance)
}
