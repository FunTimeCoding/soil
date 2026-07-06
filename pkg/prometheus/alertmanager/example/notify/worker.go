package notify

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/sound"
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
			fmt.Println("Stopped")

			return
		default:
			alerts, _ := c.MustAlerts(advanced_option.New(), nil)
			now := alerts
			add, stay, remove := difference(s.Alerts, now)
			s.Alerts = now

			for _, a := range add {
				fmt.Printf("Add: %s\n", a.Name)
			}

			if false {
				for _, a := range stay {
					fmt.Printf("Stay: %s\n", a.Name)
				}
			}

			for _, a := range remove {
				fmt.Printf("Remove: %s\n", a.Name)
			}

			if !s.Loaded {
				s.Loaded = true
			} else {
				if len(add) > 0 {
					sound.Play(sound.SosumiPath, 1.0, false)
				}

				for _, a := range add {
					var summary string

					if a.Summary == constant.None {
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
