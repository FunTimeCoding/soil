package service

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"

func New(o *option.Jellyfin) *Service {
	return &Service{source: o.Source()}
}
