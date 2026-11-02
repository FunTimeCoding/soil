package daemon

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
)

type Daemon struct {
	database      *gorm.DB
	forge         face.GitLabClient
	guard         *peer.Guard
	authorization *client.Client
	service       *service.Service
	web           *web.Server
}
