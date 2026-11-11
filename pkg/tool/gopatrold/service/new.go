package service

import "github.com/funtimecoding/soil/pkg/tool/gopatrold/types/target"

func New(t ...*target.Target) *Service {
	return &Service{targets: t}
}
