package service

import "github.com/funtimecoding/soil/pkg/tool/gosublimed/option"

func New(o *option.Sublime) *Service {
	return &Service{source: o.Source()}
}
