package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"

func New(
	j face.JiraSource,
	c face.ConfluenceSource,
) *Service {
	return &Service{jira: j, confluence: c}
}
