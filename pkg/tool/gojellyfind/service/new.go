package service

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"

func New(c face.JellyfinSource) *Service {
	return &Service{source: c}
}
