package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) Navigate(
	identifier int,
	l string,
) error {
	if l == "" {
		return validation.New(browserConstant.FirefoxLocatorRequired)
	}

	return s.client.Navigate(identifier, l)
}
