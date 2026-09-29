package service

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"

func New(c face.LinkAceSource) *Service {
	return &Service{client: c}
}
