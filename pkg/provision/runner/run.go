package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
	"time"
)

func (r *Runner) run() {
	defer r.drainChannels()

	if r.setupFunction != nil && !r.setupFunction() {
		return
	}

	r.recovery.Run(r.gitClone)

	if r.initFunction != nil {
		r.recovery.Run(r.initFunction)
	}

	r.recovery.Run(
		func() {
			r.apply(nil, constant.RunnerTriggerTimer)
		},
	)
	syncTicker := time.NewTicker(constant.RunnerSyncInterval)
	defer syncTicker.Stop()
	applyTicker := time.NewTicker(constant.RunnerApplyInterval)
	defer applyTicker.Stop()
	cleanupTicker := time.NewTicker(24 * time.Hour)
	defer cleanupTicker.Stop()

	for {
		select {
		case <-r.stop:
			return
		case <-syncTicker.C:
			if r.syncTick() {
				if r.initFunction != nil {
					r.recovery.Run(r.initFunction)
				}

				r.recovery.Run(
					func() {
						r.apply(nil, constant.RunnerTriggerTimer)
					},
				)
				applyTicker.Reset(constant.RunnerApplyInterval)
			}
		case <-applyTicker.C:
			if r.initFunction != nil {
				r.recovery.Run(r.initFunction)
			}

			r.recovery.Run(func() { r.apply(nil, constant.RunnerTriggerTimer) })
		case request := <-r.sync:
			var result *update.Result
			r.recovery.Run(func() { result = r.syncWithDiff() })

			if result == nil {
				result = update.NewResult()
				result.Error = fmt.Errorf("sync failed")
				r.syncFailures++
				r.recovery.Run(r.healRepository)
			} else {
				r.syncFailures = 0
			}

			request.Response <- result
		case request := <-r.trigger:
			if request.Update {
				r.recovery.Run(func() { r.gitSync() })
			}

			if r.initFunction != nil {
				r.recovery.Run(r.initFunction)
			}

			var value any
			r.recovery.Run(
				func() {
					value = r.apply(
						request.Parameters,
						constant.RunnerTriggerManual,
					)
				},
			)

			if request.Response != nil {
				result := trigger.NewResult()
				result.Value = value

				if value == nil {
					result.Error = fmt.Errorf("apply failed")
				}

				request.Response <- result
			}
		case <-cleanupTicker.C:
			if r.cleanupFunction != nil {
				r.recovery.Run(r.cleanupFunction)
			}
		}
	}
}
