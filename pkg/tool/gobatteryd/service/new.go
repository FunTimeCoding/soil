package service

import "github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/target"

func New(targets []*target.Target) *Service {
	return &Service{targets: targets}
}
