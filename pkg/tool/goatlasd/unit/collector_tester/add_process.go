package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func AddProcess(
	f *mock_client.Client,
	name string,
	running bool,
) {
	f.ProcessList = append(
		f.ProcessList,
		client.Process{Name: name, Command: name, Running: running},
	)
}
