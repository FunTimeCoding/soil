package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"

func New(c face.ConfluenceSource) *Service {
	return &Service{confluence: c}
}
