package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	client   *gitlab.Client
	worker   *worker.Worker
	view     *view.View
	registry *palette.Registry
}
