package service

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"

func New(o *option.Opnsense) *Service {
	return &Service{source: o.Source()}
}
