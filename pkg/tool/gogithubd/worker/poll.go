package worker

import (
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (w *Worker) Poll() {
	latest := make(map[[3]string]*run.Run)

	for _, repo := range w.client.MustRepositories(w.owner) {
		name := repo.GetName()

		for _, r := range w.client.MustLatestRuns(w.owner, name) {
			if r.Status != constant.CompletedStatus {
				continue
			}

			k := [3]string{name, r.Name, r.Branch}

			if existing, okay := latest[k]; !okay || r.Identifier > existing.Identifier {
				latest[k] = r
			}
		}
	}

	w.gauge.Reset()

	for k, r := range latest {
		repository, workflow, branch := k[0], k[1], k[2]
		w.gauge.WithLabelValues(
			w.owner,
			repository,
			workflow,
			branch,
			r.Conclusion,
		).Set(1)
	}
}
