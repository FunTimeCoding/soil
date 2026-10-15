package runner

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
)

func (r *Runner) failInventory(
	e error,
	triggerSource string,
	head string,
) []*run.Run {
	r.reporter.CaptureException(e)
	record := r.store.NewRun()
	record.Scope = constant.InventoryScope
	record.TriggerSource = triggerSource
	record.Status = provisionConstant.StoreStatusError
	record.GitHead = head
	record.ErrorOutput = e.Error()
	r.store.Create(record)

	return []*run.Run{record}
}
