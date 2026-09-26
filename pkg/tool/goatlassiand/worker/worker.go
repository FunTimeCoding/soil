package worker

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"sync"
	"time"
)

type Worker struct {
	client        *jira.Client
	confluence    *confluence.Client
	interval      time.Duration
	recovery      *recovery.Recovery
	notifier      *notifier.Notifier
	stop          chan struct{}
	mutex         sync.RWMutex
	project       []string
	issues        []*issue.Issue
	watchedIssues []*issue.Issue
	newest        []*issue.Issue
	favorites     []*page.Page
	watched       []*page.Page
}
