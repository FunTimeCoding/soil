package service

import "github.com/funtimecoding/soil/pkg/tool/gogithubd/option"

func New(o *option.Exporter) *Service {
	return &Service{source: o.Source(), owner: o.Owner}
}
