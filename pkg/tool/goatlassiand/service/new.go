package service

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/plate"
)

func New(
	j atlassian.JiraClient,
	c atlassian.ConfluenceClient,
	jiraGuard *peer.Guard,
	confluenceGuard *peer.Guard,
	project []string,
	n face.EventNotifier,
) *Service {
	return &Service{
		jira:            j,
		confluence:      c,
		jiraGuard:       jiraGuard,
		confluenceGuard: confluenceGuard,
		project:         project,
		notifier:        n,
		current:         plate.New(),
	}
}
