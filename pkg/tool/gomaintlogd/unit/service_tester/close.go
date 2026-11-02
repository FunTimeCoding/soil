package service_tester

import "github.com/funtimecoding/soil/pkg/relational"

func (o *Tester) Close() {
	relational.CloseMapper(o.database)
}
