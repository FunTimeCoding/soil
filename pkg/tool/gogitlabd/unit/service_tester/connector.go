package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"

func (o *Tester) Connector() *mock_client.Client {
	return o.connector
}
