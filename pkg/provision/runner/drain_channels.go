package runner

import "fmt"

func (r *Runner) drainChannels() {
	for {
		select {
		case request := <-r.sync:
			request.Response <- &SyncResult{Error: fmt.Errorf("runner stopped")}
		case request := <-r.trigger:
			if request.Response != nil {
				request.Response <- &TriggerResult{
					Error: fmt.Errorf("runner stopped"),
				}
			}
		default:
			return
		}
	}
}
