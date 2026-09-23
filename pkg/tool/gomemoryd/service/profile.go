package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"slices"
	"strings"
)

func (s *Service) Profile(
	topic string,
	scope string,
	detail bool,
) (*ProfileResult, *ProfileDetail, error) {
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

	result := &ProfileResult{Always: always}
	alwaysIDs := map[int64]bool{}

	for _, m := range always {
		alwaysIDs[m.Identifier] = true
	}

	allMemories, f := s.ListMemories("", "", scope, true)

	if f != nil {
		return nil, nil, fmt.Errorf("list memories: %w", f)
	}

	childSummaries := map[int64][]store.MemorySummary{}

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
				a store.MemorySummary,
				b store.MemorySummary,
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
		result.Index = append(result.Index, m)
	}

	completionsTrimmed := 0
	completionTokens := 0

	if scope == "" {
		completions, f := s.ListCompletions()

		if f == nil {
			for _, r := range completions {
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
				result.Completions = append(
					result.Completions,
					CompletionEntry{SessionName: name, Body: r.Body},
				)
			}
		}
	}

	impressionsTrimmed := 0
	impressionTokens := 0

	if scope == "" {
		impressions, f := s.LatestImpressions(10)

		if f == nil {
			for i := range impressions {
				tokens := s.tokenizer.Count(format.Impression(&impressions[i]))

				if remaining-tokens < 0 {
					impressionsTrimmed++

					continue
				}

				remaining -= tokens
				impressionTokens += tokens
				result.Impressions = append(result.Impressions, impressions[i])
			}
		}
	}

	relevantTrimmed := 0
	relevantTokens := 0

	if topic != "" && scope == "" {
		exclude := make([]string, len(always))

		for i, m := range always {
			exclude[i] = fmt.Sprintf("memory/%d", m.Identifier)
		}

		relevant, f := s.SearchRelevant(topic, 20, exclude)

		if f != nil {
			return nil, nil, fmt.Errorf("search relevant memories: %w", f)
		}

		for i := range relevant {
			tokens := s.tokenizer.Count(format.RelevantMemory(&relevant[i]))

			if remaining-tokens < 0 {
				relevantTrimmed++

				continue
			}

			remaining -= tokens
			relevantTokens += tokens
			result.Relevant = append(result.Relevant, relevant[i])
		}
	}

	result.Text = ProfileText(result)
	var d *ProfileDetail

	if detail {
		d = &ProfileDetail{
			Budget:             constant.ProfileBudget,
			AlwaysTokens:       alwaysTokens,
			IndexTokens:        indexTokens,
			IndexTrimmed:       indexTrimmed,
			CompletionTokens:   completionTokens,
			CompletionsTrimmed: completionsTrimmed,
			ImpressionTokens:   impressionTokens,
			ImpressionsTrimmed: impressionsTrimmed,
			RelevantTokens:     relevantTokens,
			RelevantTrimmed:    relevantTrimmed,
			TotalTokens:        constant.ProfileBudget - remaining,
		}
	}

	return result, d, nil
}
