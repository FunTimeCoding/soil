package collector_tester

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"

func (o *Tester) LeaseClient() *mock_client.Client {
	o.t.Helper()

	return mock_client.New(o.t)
}
