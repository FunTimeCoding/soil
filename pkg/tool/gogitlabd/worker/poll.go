package worker

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/latest_pipeline"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (w *Worker) Poll() {
	type key struct {
		project   string
		reference string
	}
	latest := make(map[key]*gitlab.PipelineInfo)

	for _, p := range w.client.MustProjects() {
		project := fmt.Sprintf("%s/%s", p.Namespace, p.Name)

		for reference, i := range latest_pipeline.New(
			w.client.MustBranches(p.Identifier),
			w.client.MustPipelines(p.Identifier),
		) {
			latest[key{project, reference}] = i
		}
	}

	w.gauge.Reset()

	for k, i := range latest {
		w.gauge.WithLabelValues(k.project, k.reference, i.Status).Set(1)
	}
}
