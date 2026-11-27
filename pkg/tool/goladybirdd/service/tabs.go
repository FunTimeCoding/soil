package service

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"

func (s *Service) Tabs() ([]*response.Context, error) {
	return guarded(s, s.client.Tree)
}
