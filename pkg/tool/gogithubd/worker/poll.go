package worker

import "github.com/funtimecoding/soil/pkg/github/run"

func (w *Worker) Poll() {
	type key struct {
		repo     string
		workflow string
		branch   string
	}
	latest := make(map[key]*run.Run)

	for _, repo := range w.client.MustRepositories(w.owner) {
		name := repo.GetName()

		for _, r := range w.client.MustLatestRuns(w.owner, name) {
			if r.Status != run.Completed {
				continue
			}

			k := key{name, r.Name, r.Branch}

			if existing, okay := latest[k]; !okay || r.Identifier > existing.Identifier {
				latest[k] = r
			}
		}
	}

	w.gauge.Reset()

	for k, r := range latest {
		w.gauge.WithLabelValues(
			w.owner,
			k.repo,
			k.workflow,
			k.branch,
			r.Conclusion,
		).Set(1)
	}
}
