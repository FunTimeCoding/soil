package worker

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func (w *Worker) RecoverStale() {
	alerts, _ := w.client.MustAlerts(advanced_option.New(), nil)
	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State == constant.ActiveState {
			current[a.Fingerprint] = true
		}
	}

	unresolved := w.store.MustUnresolved()

	for _, u := range unresolved {
		if current[u.Fingerprint] {
			w.firing[u.Fingerprint] = u.Key
		}
	}

	var resolved int

	for _, u := range unresolved {
		if w.firing[u.Fingerprint] == u.Key {
			continue
		}

		w.store.MustResolve(u.Key)
		resolved++
	}

	w.logger.Structured(
		"recovered stale records",
		"adopted",
		len(w.firing),
		"resolved",
		resolved,
	)
}
