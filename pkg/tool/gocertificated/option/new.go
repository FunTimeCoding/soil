package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Certificate {
	return &Certificate{
		Web:   option.Web{Authorization: client.NewEnvironment},
		Forge: func() face.GitLabClient { return gitlab.NewEnvironment() },
	}
}
