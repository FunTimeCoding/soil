package daemon

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"sync"
)

type Daemon struct {
	jira            atlassian.JiraClient
	confluence      atlassian.ConfluenceClient
	jiraGuard       *peer.Guard
	confluenceGuard *peer.Guard
	service         *service.Service
	web             *web.Server
	logger          *logger.Logger
	reporter        face.Reporter
	closeOnce       sync.Once
}
