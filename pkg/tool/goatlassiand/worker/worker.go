package worker

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"sync"
	"time"
)

type Worker struct {
	service       *service.Service
	interval      time.Duration
	logger        *logger.Logger
	reporter      face.Reporter
	recovery      *recovery.Recovery
	notifier      *notifier.Notifier
	stop          chan struct{}
	once          sync.Once
	mutex         sync.RWMutex
	issues        []*issue.Issue
	watchedIssues []*issue.Issue
	newest        []*issue.Issue
	favorites     []*page.Page
	watched       []*page.Page
}
