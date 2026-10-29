package service

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/option"

func New(o *option.Secret) *Service {
	return &Service{source: o.Source()}
}
