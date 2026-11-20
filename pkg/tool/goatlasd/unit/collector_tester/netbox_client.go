package collector_tester

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"

func (o *Tester) NetboxClient() *mock_client.Client {
	o.t.Helper()

	return mock_client.New(o.t)
}
