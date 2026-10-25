package example

import (
	domain "example/pkg/domain/face"
	"example/pkg/face"
	demo "example/pkg/tool/demo/face"
)

type Service struct{}

type Concrete struct{}

type reporter struct{}

func (reporter) Report() {}

type client struct{}

func (client) Fetch() {}

func NewService(
	r face.Reporter,
	c demo.Client,
	t domain.Thing,
	more ...face.Reporter,
) *Service {
	return &Service{}
}

func NewHolder(c *Concrete) *Service {
	return &Service{}
}

func Build(r face.Reporter) *Service {
	return &Service{}
}
