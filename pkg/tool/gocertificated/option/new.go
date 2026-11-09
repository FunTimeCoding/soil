package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

func New() *Certificate {
	return &Certificate{
		Web:   option.Web{Authorization: client.NewEnvironment},
		Clock: time.Now,
		Forge: func() face.GitLabClient { return gitlab.NewEnvironment() },
	}
}
