package service_tester

import "github.com/funtimecoding/soil/pkg/relational"

func (o *Tester) CloseStore() {
	relational.CloseMapper(o.mapper)
}
