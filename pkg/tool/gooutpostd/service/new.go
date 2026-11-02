package service

import "github.com/funtimecoding/soil/pkg/system/service"

func New(reader *service.Client) *Service {
	return &Service{reader: reader}
}
