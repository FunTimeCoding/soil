package request_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
)

func (s *Server) List(f *merge_request_filter.Filter) []*merge_request.Request {
	s.t.Helper()
	result, e := s.client.MergeRequests(s.t.Context(), f)
	assert.FatalOnError(s.t, e)

	return result
}
