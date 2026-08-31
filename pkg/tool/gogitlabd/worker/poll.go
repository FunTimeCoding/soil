package worker

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/latest_pipeline"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"slices"
	"strings"
	"time"
)

func (w *Worker) Poll() {
	type key struct {
		project   string
		reference string
	}
	latest := make(map[key]*gitlab.PipelineInfo)
	projects := make(map[string]*project.Project)

	for _, p := range w.client.MustProjects() {
		name := join.Slash([]string{p.Namespace, p.Name})
		projects[name] = p

		for reference, i := range latest_pipeline.New(
			w.client.MustBranches(p.Identifier),
			w.client.MustPipelines(p.Identifier),
		) {
			latest[key{name, reference}] = i
		}
	}

	w.gauge.Reset()
	entries := make([]*board_entry.Entry, 0, len(latest))

	for k, i := range latest {
		w.gauge.WithLabelValues(k.project, k.reference, i.Status).Set(1)
		var updated time.Time

		if i.UpdatedAt != nil {
			updated = *i.UpdatedAt
		}

		entries = append(
			entries,
			board_entry.New(
				k.project,
				projects[k.project].Identifier,
				projects[k.project].Raw.WebURL,
				k.reference,
				i.Status,
				i.ID,
				i.WebURL,
				updated,
			),
		)
	}

	slices.SortFunc(
		entries,
		func(a *board_entry.Entry, b *board_entry.Entry) int {
			if c := b.Updated.Compare(a.Updated); c != 0 {
				return c
			}

			if c := strings.Compare(a.Project, b.Project); c != 0 {
				return c
			}

			return strings.Compare(a.Reference, b.Reference)
		},
	)
	w.mutex.Lock()
	previous := w.entries
	w.entries = entries
	w.mutex.Unlock()

	if !slices.EqualFunc(
		previous,
		entries,
		func(a *board_entry.Entry, b *board_entry.Entry) bool {
			return *a == *b
		},
	) {
		w.notifier.Notify()
	}
}
