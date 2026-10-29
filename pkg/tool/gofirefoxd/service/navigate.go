package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
)

func (s *Service) Navigate(
	identifier int,
	l string,
) error {
	if l == "" {
		return validation.New(constant.LocatorRequired)
	}

	return s.client.Navigate(identifier, l)
}
