package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"

func (o *Tester) Catalog() *catalog.Service {
	o.t.Helper()

	if o.catalog == nil {
		o.catalog = catalog.New(o.peers())
	}

	return o.catalog
}
