package service

import "github.com/funtimecoding/soil/pkg/firefox/content"

func (s *Service) ReadTab(
	identifier int,
	raw bool,
) (*content.Content, error) {
	return s.client.ReadTab(identifier, raw)
}
