package service

import (
	confluence "github.com/funtimecoding/soil/pkg/atlassian/confluence/face"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	jira            jira.Client
	confluence      confluence.Client
	jiraGuard       *peer.Guard
	confluenceGuard *peer.Guard
	project         []string
}
