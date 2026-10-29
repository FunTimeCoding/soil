package service

import "github.com/funtimecoding/soil/pkg/system/service"

func New() *Service {
	return &Service{reader: service.New()}
}
