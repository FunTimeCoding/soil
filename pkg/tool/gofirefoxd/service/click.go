package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/firefox/element"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) Click(
	identifier int,
	selector string,
) (*element.Element, error) {
	if selector == "" {
		return nil, validation.New(browserConstant.FirefoxSelectorRequired)
	}

	return s.client.Click(identifier, selector)
}
