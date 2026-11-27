package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/firefox/element"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) Fill(
	identifier int,
	selector string,
	value string,
	direct bool,
) (*element.Element, error) {
	if selector == "" {
		return nil, validation.New(browserConstant.FirefoxSelectorRequired)
	}

	return s.client.Fill(identifier, selector, value, direct)
}
