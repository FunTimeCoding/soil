package worker

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"time"
)

func (w *Worker) Poll() {
	start := time.Now()
	w.lastPoll.Store(start)

	if w.metrics != nil {
		w.metrics.lastPollTime.SetToCurrentTime()
	}

	defer func() {
		if w.metrics != nil {
			w.metrics.pollDuration.Observe(
				time.Since(start).Seconds(),
			)
		}
	}()
	alerts, _ := w.client.MustAlerts(advanced_option.New(), nil)
	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State != constant.ActiveState {
			continue
		}

		current[a.Fingerprint] = true

		if _, exists := w.firing[a.Fingerprint]; exists {
			continue
		}

		labels := make(map[string]string)

		for k, v := range a.Labels {
			labels[k] = v
		}

		r := store.Record{
			Fingerprint: a.Fingerprint,
			Name:        a.Name,
			Severity:    a.Severity,
			Summary:     a.Summary,
			Labels:      labels,
			Start:       time.Now(),
		}
		w.firing[a.Fingerprint] = w.store.MustSave(r)

		if w.metrics != nil {
			w.metrics.alertsTotal.Inc()
		}
	}

	for fingerprint, key := range w.firing {
		if current[fingerprint] {
			continue
		}

		w.store.MustResolve(key)
		delete(w.firing, fingerprint)
	}

	if pruned := w.store.MustPrune(time.Now().Add(-w.retention)); pruned > 0 {
		w.logger.Structured("pruned records", "count", pruned)
	}

	if w.metrics != nil {
		w.metrics.alertsFiring.Set(float64(len(w.firing)))
		w.metrics.recordsTotal.Set(float64(w.store.MustCount()))
	}
}
