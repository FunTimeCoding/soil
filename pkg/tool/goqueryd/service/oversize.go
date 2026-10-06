package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize"
	"sort"
)

func (s *Service) Oversize(collection string) (*oversize.Report, error) {
	documents, e := s.store.DocumentBodies(collection)

	if e != nil {
		return oversize.Stub(), e
	}

	result := oversize.New(
		s.reranker.Name(),
		s.reranker.SequenceLength(),
		s.reranker.Allowance(),
	)

	for _, d := range documents {
		if f := s.oversizeFile(d, result.Allowance); f != nil {
			result.Files = append(result.Files, f)
		}
	}

	sort.Slice(
		result.Files,
		func(i, j int) bool {
			a, b := result.Files[i], result.Files[j]

			if a.Worst != b.Worst {
				return a.Worst > b.Worst
			}

			if a.Collection != b.Collection {
				return a.Collection < b.Collection
			}

			return a.Path < b.Path
		},
	)

	return result, nil
}
