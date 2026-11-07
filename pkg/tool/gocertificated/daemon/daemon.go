package daemon

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"gorm.io/gorm"
)

type Daemon struct {
	database *gorm.DB
	forge    face.GitLabClient
	guard    *peer.Guard
	service  *service.Service
	web      *web.Server
}
