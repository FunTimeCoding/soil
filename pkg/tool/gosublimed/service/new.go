package service

import "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"

func New(c face.SublimeSource) *Service {
	return &Service{source: c}
}
