package worker

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
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
	open := make(map[string]bool)

	for _, fingerprint := range w.store.MustUnresolvedFingerprints() {
		open[fingerprint] = true
	}

	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State != constant.ActiveState {
			continue
		}

		current[a.Fingerprint] = true

		if open[a.Fingerprint] {
			continue
		}

		labels := make(map[string]string)

		for k, v := range a.Labels {
			labels[k] = v
		}

		w.store.MustCreate(
			record.Record{
				Fingerprint: a.Fingerprint,
				Name:        a.Name,
				Severity:    a.Severity,
				Summary:     a.Summary,
				Labels:      labels,
				Start:       time.Now(),
			},
		)
		open[a.Fingerprint] = true

		if w.metrics != nil {
			w.metrics.alertsTotal.Inc()
		}
	}

	for fingerprint := range open {
		if current[fingerprint] {
			continue
		}

		w.store.MustResolve(fingerprint)
	}

	if pruned := w.store.MustPrune(time.Now().Add(-w.retention)); pruned > 0 {
		w.logger.Structured("pruned records", "count", pruned)
	}

	if w.metrics != nil {
		w.metrics.alertsFiring.Set(float64(w.store.MustUnresolvedCount()))
		w.metrics.recordsTotal.Set(float64(w.store.MustCount()))
	}
}
