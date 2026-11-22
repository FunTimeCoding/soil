package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.nationalGuard, d.trackerGuard, d.extendedGuard}

	if d.ansibleGuard != nil {
		result = append(result, d.ansibleGuard)
	}

	result = append(result, d.clock, d.scoreClock)

	if d.feedClock != nil {
		result = append(result, d.feedClock)
	}

	return result
}
