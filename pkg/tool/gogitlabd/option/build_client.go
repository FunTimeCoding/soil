package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web"
)

func BuildClient(i *instance.Instance) (face.GitLabClient, error) {
	return gitlab.New(
		web.TrimScheme(i.Host),
		i.Token,
		gitlab.WithGroups(i.Groups),
		gitlab.WithProjects(i.Projects),
		gitlab.WithRetryMax(1),
	), nil
}
