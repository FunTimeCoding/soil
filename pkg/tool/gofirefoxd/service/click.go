package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (s *Service) Click(
	identifier int,
	selector string,
) (*element.Element, error) {
	if selector == "" {
		return nil, validation.New(constant.SelectorRequired)
	}

	return s.client.Click(identifier, selector)
}
