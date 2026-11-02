package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
)

func (o *Tester) StoredRuns() []run.Run {
	o.t.Helper()
	d := lite.New(logger.New(o.t.Context()), o.litePath)
	defer relational.CloseMapper(d)
	result, e := store.New(d, constant.RunsTable).Recent(10)

	if e != nil {
		o.t.Fatal(e)
	}

	return result
}
