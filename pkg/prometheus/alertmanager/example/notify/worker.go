package notify

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/sound"
	soundConstant "github.com/funtimecoding/soil/pkg/sound/constant"
	"github.com/funtimecoding/soil/pkg/system/macos"
	"time"
)

func worker(
	stop <-chan struct{},
	c *alertmanager.Client,
	s *State,
) {
	for {
		select {
		case <-stop:
			console.Line("Stopped")

			return
		default:
			alerts, _ := c.MustAlerts(advanced_option.New(), nil)
			now := alerts
			add, stay, remove := difference(s.Alerts, now)
			s.Alerts = now

			for _, a := range add {
				console.Format("Add: %s\n", a.Name)
			}

			if false {
				for _, a := range stay {
					console.Format("Stay: %s\n", a.Name)
				}
			}

			for _, a := range remove {
				console.Format("Remove: %s\n", a.Name)
			}

			if !s.Loaded {
				s.Loaded = true
			} else {
				if len(add) > 0 {
					sound.Play(soundConstant.SosumiPath, 1.0, false)
				}

				for _, a := range add {
					var summary string

					if a.Summary == prometheus.None {
						summary = "no summary"
					} else {
						summary = a.Summary
					}

					macos.Notify("Alert", a.Name, summary)
				}
			}

			time.Sleep(10 * time.Second)
		}
	}
}
