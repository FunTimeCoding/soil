package hub

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/face"
)

func Run(
	c face.Client,
	namespace string,
	repository string,
) {
	for _, r := range c.MustProjectRuns(namespace, repository) {
		if r.Status != constant.CompletedStatus {
			continue
		}

		console.Format("Delete run: %d\n", r.Identifier)
		c.MustDeleteRun(namespace, repository, r.Identifier)
	}
}
