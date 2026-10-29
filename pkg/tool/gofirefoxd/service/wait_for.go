package service

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (s *Service) WaitFor(
	identifier int,
	selector string,
	timeout int,
) (*tab.Tab, error) {
	return s.client.WaitFor(identifier, selector, timeout)
}
