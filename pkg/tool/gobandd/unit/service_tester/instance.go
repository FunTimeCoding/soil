package service_tester

import "github.com/funtimecoding/soil/pkg/band/mock_client"

func (o *Tester) Instance(name string) *mock_client.Client {
	o.t.Helper()
	result, okay := o.instances[name]

	if !okay {
		o.t.Fatalf("no instance %s", name)
	}

	return result
}
