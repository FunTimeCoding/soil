package service_tester

import "github.com/funtimecoding/soil/pkg/band/mock_client"

func (o *Tester) Bravo() *mock_client.Client {
	o.t.Helper()

	return o.Instance("bravo")
}
