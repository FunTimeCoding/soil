package worker

import "github.com/funtimecoding/soil/pkg/smartctl/response"

func (w *Worker) drive(
	host string,
	r *response.Report,
) {
	labels := []string{host, r.Device.Name, r.ModelName}

	if r.Status != nil {
		w.passed.WithLabelValues(labels...).Set(flag(r.Status.Passed))
	}

	if r.Temperature != nil {
		w.temperature.WithLabelValues(labels...).Set(
			float64(r.Temperature.Current),
		)
	}

	if h := r.Health; h != nil {
		w.warning.WithLabelValues(labels...).Set(float64(h.CriticalWarning))
		w.spare.WithLabelValues(labels...).Set(float64(h.AvailableSpare))
		w.used.WithLabelValues(labels...).Set(float64(h.PercentageUsed))
		w.media.WithLabelValues(labels...).Set(float64(h.MediaErrors))
	}
}
