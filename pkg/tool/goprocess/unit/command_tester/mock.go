package command_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (o *Tester) Mock() *mock_client.Client {
	o.t.Helper()
	result := mock_client.New(o.t)
	result.ProcessList = []client.Process{
		{Name: "alfa", Running: true},
		{Name: "bravo"},
	}

	return result
}
