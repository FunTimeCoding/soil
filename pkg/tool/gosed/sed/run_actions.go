package sed

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/strings/base64"
	"github.com/funtimecoding/soil/pkg/tool/gosed/sed/option"
	upstream "gitlab.com/gitlab-org/api/client-go/v2"
	"log"
)

func runActions(
	c *gitlab.Client,
	project int64,
	o *option.Sed,
) {
	actions := make([]*Action, len(o.RawActions))

	for i, raw := range o.RawActions {
		actions[i] = ParseAction(raw)
	}

	grouped := make(map[string][]*Action)
	var order []string

	for _, a := range actions {
		if _, exists := grouped[a.Path]; !exists {
			order = append(order, a.Path)
		}

		grouped[a.Path] = append(grouped[a.Path], a)
	}

	var commitActions []*upstream.CommitActionOptions
	update := upstream.FileUpdate

	for _, path := range order {
		f := c.MustFile(project, o.Branch, path)

		if f == nil {
			log.Panicf("file does not exist: %s", path)
		}

		replaces := make(map[string]string)

		for _, a := range grouped[path] {
			replaces[a.Prefix] = a.Value
		}

		commitActions = append(
			commitActions,
			&upstream.CommitActionOptions{
				Action:   &update,
				FilePath: new(path),
				Content: new(
					replaceLines(base64.Decode(f.Content), replaces),
				),
			},
		)
	}

	c.MustCommitActions(project, o.Branch, o.Message, commitActions)
}
