package web_service_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"

func (o *Tester) Store() *store.Store {
	return o.server.Store
}
