package collector_tester

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"

func (o *Tester) ProcessClient() *mock_client.Client {
	o.t.Helper()

	return mock_client.New(o.t)
}
