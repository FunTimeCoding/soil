package worker

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/latest_pipeline"
	"time"
)

func (w *Worker) pollInstance(instance string) ([]*board_entry.Entry, error) {
	latest := make(map[[2]string]*pipeline.Pipeline)
	projects := make(map[string]*project.Project)
	all, e := w.service.Projects(instance, "")

	if e != nil {
		return nil, e
	}

	for _, p := range all {
		name := join.Slash([]string{p.Namespace, p.Name})
		projects[name] = p
		branches, f := w.service.Branches(instance, p.Identifier)

		if f != nil {
			return nil, f
		}

		tags, g := w.service.Tags(instance, p.Identifier)

		if g != nil {
			return nil, g
		}

		pipelines, h := w.service.Pipelines(instance, p.Identifier, "", "", 0)

		if h != nil {
			return nil, h
		}

		for reference, i := range latest_pipeline.New(branches, tags, pipelines) {
			latest[[2]string{name, reference}] = i
		}
	}

	result := make([]*board_entry.Entry, 0, len(latest))

	for k, i := range latest {
		name, reference := k[0], k[1]
		var updated time.Time

		if i.Update != nil {
			updated = *i.Update
		}

		result = append(
			result,
			board_entry.New(
				instance,
				name,
				projects[name].Identifier,
				projects[name].Raw.WebURL,
				reference,
				i.Status,
				i.Identifier,
				i.Link,
				updated,
			),
		)
	}

	return result, nil
}
