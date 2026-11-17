package service_tester

import "github.com/funtimecoding/soil/pkg/ssh/mock_client"

func (o *Tester) Dialled() []*mock_client.Client {
	return o.dialled
}
