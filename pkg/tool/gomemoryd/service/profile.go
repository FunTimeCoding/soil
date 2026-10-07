package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile"
	"slices"
	"strings"
)

func (s *Service) Profile(
	topic string,
	scope string,
	detail bool,
) (*profile.Result, *profile.Detail, error) {
	if scope == constant.AllScope || scope == constant.DefaultScope {
		return nil, nil, validation.New("scope name is reserved: %s", scope)
	}

	if f := s.ensureTokenizer(); f != nil {
		return nil, nil, f
	}

	always, e := s.ListMemoriesWithContent(constant.AlwaysTag, scope)

	if e != nil {
		return nil, nil, fmt.Errorf("load always memories: %w", e)
	}

	alwaysIDs := map[int64]bool{}

	for _, m := range always {
		alwaysIDs[m.Identifier] = true
	}

	allMemories, f := s.ListMemories("", "", scope, true)

	if f != nil {
		return nil, nil, fmt.Errorf("list memories: %w", f)
	}

	childSummaries := map[int64][]record.MemorySummary{}

	for _, m := range allMemories {
		if m.ParentIdentifier != nil {
			childSummaries[*m.ParentIdentifier] = append(
				childSummaries[*m.ParentIdentifier],
				m,
			)
		}
	}

	childrenByParent := map[int64][]string{}

	for parent, children := range childSummaries {
		slices.SortStableFunc(
			children,
			func(
				a record.MemorySummary,
				b record.MemorySummary,
			) int {
				return a.Ordinal - b.Ordinal
			},
		)
		names := make([]string, 0, len(children))

		for _, m := range children {
			names = append(names, m.Name)
		}

		childrenByParent[parent] = names
	}

	alwaysTokens := 0

	for i := range always {
		if children, found := childrenByParent[always[i].Identifier]; found {
			always[i].Children = children
		}

		alwaysTokens += s.tokenizer.Count(format.AlwaysMemory(&always[i]))
	}

	remaining := max(constant.ProfileBudget-alwaysTokens, 0)
	indexTrimmed := 0
	indexTokens := 0
	var index []record.MemorySummary

	for _, m := range allMemories {
		if alwaysIDs[m.Identifier] {
			continue
		}

		if m.ParentIdentifier != nil {
			continue
		}

		if slices.Contains(m.Tags, constant.NoIndexTag) {
			continue
		}

		if children, found := childrenByParent[m.Identifier]; found {
			m.Children = children
		}

		tokens := s.tokenizer.Count(format.IndexEntry(&m))

		if remaining-tokens < 0 {
			indexTrimmed++

			continue
		}

		remaining -= tokens
		indexTokens += tokens
		index = append(index, m)
	}

	completionsTrimmed := 0
	completionTokens := 0
	var completions []*profile.Completion

	if scope == "" {
		results, g := s.ListCompletions()

		if g == nil {
			for _, r := range results {
				name := r.Path

				if i := strings.LastIndex(name, stringConstant.Slash); i >= 0 {
					name = name[:i]
				}

				tokens := s.tokenizer.Count(format.Completion(name, r.Body))

				if remaining-tokens < 0 {
					completionsTrimmed++

					continue
				}

				remaining -= tokens
				completionTokens += tokens
				completions = append(
					completions,
					profile.NewCompletion(name, r.Body),
				)
			}
		}
	}

	impressionsTrimmed := 0
	impressionTokens := 0
	var impressions []record.Impression

	if scope == "" {
		latest, g := s.LatestImpressions(10)

		if g == nil {
			for i := range latest {
				tokens := s.tokenizer.Count(format.Impression(&latest[i]))

				if remaining-tokens < 0 {
					impressionsTrimmed++

					continue
				}

				remaining -= tokens
				impressionTokens += tokens
				impressions = append(impressions, latest[i])
			}
		}
	}

	relevantTrimmed := 0
	relevantTokens := 0
	var relevant []record.SearchResult

	if topic != "" && scope == "" {
		exclude := make([]string, len(always))

		for i, m := range always {
			exclude[i] = fmt.Sprintf("memory/%d", m.Identifier)
		}

		found, g := s.SearchRelevant(topic, 20, exclude)

		if g != nil {
			return nil, nil, fmt.Errorf("search relevant memories: %w", g)
		}

		for i := range found {
			tokens := s.tokenizer.Count(format.RelevantMemory(&found[i]))

			if remaining-tokens < 0 {
				relevantTrimmed++

				continue
			}

			remaining -= tokens
			relevantTokens += tokens
			relevant = append(relevant, found[i])
		}
	}

	result := profile.NewResult(
		always,
		index,
		relevant,
		impressions,
		completions,
	)

	if !detail {
		return result, nil, nil
	}

	d := profile.NewDetail()
	d.Budget = constant.ProfileBudget
	d.AlwaysTokens = alwaysTokens
	d.IndexTokens = indexTokens
	d.IndexTrimmed = indexTrimmed
	d.CompletionTokens = completionTokens
	d.CompletionsTrimmed = completionsTrimmed
	d.ImpressionTokens = impressionTokens
	d.ImpressionsTrimmed = impressionsTrimmed
	d.RelevantTokens = relevantTokens
	d.RelevantTrimmed = relevantTrimmed
	d.TotalTokens = constant.ProfileBudget - remaining

	return result, d, nil
}
