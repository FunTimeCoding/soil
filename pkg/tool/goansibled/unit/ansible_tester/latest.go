package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
)

func (o *Tester) latest() (*run.Run, bool) {
	d := lite.New(logger.New(o.t.Context()), o.litePath)
	defer relational.CloseMapper(d)
	runs, e := store.New(d, constant.RunsTable).Recent(1)

	if e != nil || len(runs) != 1 {
		return nil, false
	}

	return &runs[0], true
}
