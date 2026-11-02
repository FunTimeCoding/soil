package service

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"

func New(s face.SecretSource) *Service {
	return &Service{source: s}
}
