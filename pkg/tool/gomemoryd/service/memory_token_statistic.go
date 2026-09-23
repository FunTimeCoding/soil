package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"

func (s *Service) MemoryTokenStatistic(
	identifier int64,
	scope string,
) (*token_summary.Statistic, *token_summary.Summary, error) {
	summary, e := s.MemoryTokenSummary(scope)

	if e != nil {
		return nil, nil, e
	}

	for _, t := range summary.Statistic {
		if t.Identifier == identifier {
			return t, summary, nil
		}
	}

	return nil, summary, nil
}
