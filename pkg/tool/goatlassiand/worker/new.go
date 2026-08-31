package worker

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func New(
	client *jira.Client,
	c *confluence.Client,
	interval time.Duration,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	return &Worker{
		client:     client,
		confluence: c,
		interval:   interval,
		recovery:   recovery.New(l, r),
		notifier:   notifier.New(),
		stop:       make(chan struct{}),
	}
}
