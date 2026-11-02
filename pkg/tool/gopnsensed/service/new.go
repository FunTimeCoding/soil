package service

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"

func New(c face.OpnsenseSource) *Service {
	return &Service{source: c}
}
