package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
)

func (o *Tester) runs() []run.Run {
	d := lite.New(logger.New(o.t.Context()), o.litePath)
	defer relational.CloseMapper(d)
	result, e := store.New(d, constant.RunsTable).Recent(10)

	if e != nil {
		return nil
	}

	return result
}
