package service

import "github.com/funtimecoding/soil/pkg/tool/goitermd/option"

func New(o *option.Iterm) *Service {
	return &Service{source: o.Source(), keyInterval: o.KeyInterval}
}
