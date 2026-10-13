package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web"
)

func newClient(i *instance.Instance) (face.Forge, error) {
	return gitlab.New(
		web.TrimScheme(i.Host),
		i.Token,
		gitlab.WithGroups(i.Groups),
		gitlab.WithProjects(i.Projects),
	), nil
}
