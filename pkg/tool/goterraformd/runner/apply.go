package runner

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"time"
)

func (r *Runner) apply(
	parameters map[string]any,
	triggerSource string,
) any {
	r.clearOrphanedLock()
	r.pending = nil
	record := r.store.NewRun()
	record.TriggerSource = triggerSource
	record.Status = provisionConstant.StoreStatusRunning
	record.GitHead = r.provision.CurrentHead()
	r.store.Create(record)
	r.logger.Structured("terraform_apply_start")
	target, _ := parameters[constant.Target].(string)
	record.Scope = target
	start := time.Now()
	e := r.planApply(record, target)
	record.DurationMillisecond = time.Since(start).Milliseconds()

	if e != nil {
		record.Status = provisionConstant.StoreStatusError
		r.logger.Structured("terraform_apply_error", "error", e.Error())
	} else {
		record.Status = provisionConstant.StoreStatusSuccess
		r.metrics.LastSuccess.Set(float64(time.Now().Unix()))
		r.logger.Structured("terraform_apply_done")
	}

	r.metrics.RunsTotal.WithLabelValues(record.Status).Inc()
	r.metrics.ApplyDuration.Observe(time.Since(start).Seconds())
	r.store.Update(record)

	return record
}
