package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (s *Service) Fill(
	identifier int,
	selector string,
	value string,
	direct bool,
) (*element.Element, error) {
	if selector == "" {
		return nil, validation.New(constant.SelectorRequired)
	}

	return s.client.Fill(identifier, selector, value, direct)
}
