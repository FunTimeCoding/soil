package service

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/plate"
	"sync"
)

type Service struct {
	jira            atlassian.JiraClient
	confluence      atlassian.ConfluenceClient
	jiraGuard       *peer.Guard
	confluenceGuard *peer.Guard
	project         []string
	notifier        face.EventNotifier
	mutex           sync.RWMutex
	current         *plate.Plate
}
