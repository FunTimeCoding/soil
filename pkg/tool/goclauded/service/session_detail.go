package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/session_detail"

func (s *Service) SessionDetail(query string) (*session_detail.Detail, error) {
	r, e := s.store.ResolveSessionIdentifier(query)

	if e != nil {
		return nil, e
	}

	if !r.Found() {
		return nil, nil
	}

	v, f := s.store.GetSession(r.Identifier())

	if f != nil {
		return nil, f
	}

	if v == nil {
		return nil, nil
	}

	completions, f := s.store.CompletionsBySession(v.Identifier)

	if f != nil {
		return nil, f
	}

	body, f := s.store.SummaryBySession(v.Identifier)

	if f != nil {
		return nil, f
	}

	result := session_detail.New()
	result.Identifier = v.Identifier
	result.Slug = v.Slug
	result.Created = v.SessionTimestamp
	result.Session = v
	result.Alias = v.AliasValue()
	result.Description = v.Description
	result.TurnCount = v.TurnCount
	result.Completions = completions
	result.Summary = body

	return result, nil
}
