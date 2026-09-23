package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"slices"
)

func (s *Service) MemoryTokenSummary(
	scope string,
) (*token_summary.Summary, error) {
	if f := s.ensureTokenizer(); f != nil {
		return nil, f
	}

	memories, e := s.ListMemoriesWithContent("", scope)

	if e != nil {
		return nil, fmt.Errorf("load memories: %w", e)
	}

	hiddenTag := s.HiddenTag()
	statistic := make([]*token_summary.Statistic, 0, len(memories))
	blocks := make([]int, 0, len(memories))
	descriptions := make([]int, 0, len(memories))
	withheld := 0

	for i := range memories {
		block := s.tokenizer.Count(format.AlwaysMemory(&memories[i]))
		description := s.tokenizer.Count(memories[i].Description)
		hidden := hiddenTag != "" &&
			slices.Contains(memories[i].Tags, hiddenTag)

		if hidden {
			withheld++
		}

		statistic = append(
			statistic,
			token_summary.NewStatistic(
				memories[i].Identifier,
				memories[i].Name,
				memories[i].Tags,
				block,
				description,
				hidden,
			),
		)
		blocks = append(blocks, block)
		descriptions = append(descriptions, description)
	}

	slices.SortStableFunc(
		statistic,
		func(
			a *token_summary.Statistic,
			b *token_summary.Statistic,
		) int {
			return b.Block - a.Block
		},
	)

	return token_summary.New(
		statistic,
		token_summary.NewSpread(blocks),
		token_summary.NewSpread(descriptions),
		withheld,
	), nil
}
