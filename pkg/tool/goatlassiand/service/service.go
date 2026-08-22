package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"

type Service struct {
	jira       face.JiraSource
	confluence face.ConfluenceSource
}
