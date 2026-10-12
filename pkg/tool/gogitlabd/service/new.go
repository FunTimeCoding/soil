package service

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func New(c face.Forge) *Service {
	return &Service{client: c}
}
