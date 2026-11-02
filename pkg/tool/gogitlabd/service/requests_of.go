package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
)

func (s *Service) requestsOf(
	x context.Context,
	instance string,
) ([]*request_entry.Entry, error) {
	assigned, e := s.AssignedMergeRequests(x, instance, false)

	if e != nil {
		return nil, e
	}

	reviewing, f := s.ReviewingMergeRequests(x, instance, false)

	if f != nil {
		return nil, f
	}

	var result []*request_entry.Entry

	for _, r := range MergeRequests(assigned, reviewing) {
		result = append(result, request_entry.New(instance, r))
	}

	return result, nil
}
