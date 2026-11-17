package service_tester

func (o *Tester) Instance(name string) *Clients {
	o.t.Helper()
	result, okay := o.instances[name]

	if !okay {
		o.t.Fatalf("no instance %s", name)
	}

	return result
}
