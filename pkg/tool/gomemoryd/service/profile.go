package service

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
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
		return nil, nil, fmt.Errorf("%w: %s", constant.ErrorReservedScope, scope)
	}

	if f := s.ensureTokenizer(); f != nil {
		return nil, nil, f
	}

	always, e := s.ListMemoriesWithContent(constant.AlwaysTag, scope)

	if e != nil {
		return nil, nil, fmt.Errorf("%w: %v", constant.ErrorAlwaysLoad, e)
	}

	result := &ProfileResult{Always: always}
	alwaysTokens := s.countTokens(always)
	remaining := constant.ProfileBudget - alwaysTokens

	if remaining < 0 {
		remaining = 0
	}

	alwaysIDs := map[int64]bool{}

	for _, m := range always {
		alwaysIDs[m.Identifier] = true
	}

	allMemories, e := s.ListMemories("", "", scope, true)

	if e != nil {
		return nil, nil, fmt.Errorf("%w: %v", constant.ErrorMemoryList, e)
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

	indexTrimmed := 0

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

		tokens := s.countTokens(m)

		if remaining-tokens < 0 {
			indexTrimmed++

			continue
		}

		remaining -= tokens
		result.Index = append(result.Index, m)
	}

	completionsTrimmed := 0

	if scope == "" {
		completions, f := s.ListCompletions()

		if f == nil {
			for _, r := range completions {
				name := r.Path

				if i := strings.LastIndex(
					name,
					stringsConstant.Slash,
				); i >= 0 {
					name = name[:i]
				}

				entry := CompletionEntry{SessionName: name, Body: r.Body}
				tokens := s.countTokens(entry)

				if remaining-tokens < 0 {
					completionsTrimmed++

					continue
				}

				remaining -= tokens
				result.Completions = append(result.Completions, entry)
			}
		}
	}

	impressionsTrimmed := 0

	if scope == "" {
		impressions, f := s.LatestImpressions(10)

		if f == nil {
			for _, i := range impressions {
				tokens := s.countTokens(i)

				if remaining-tokens < 0 {
					impressionsTrimmed++

					continue
				}

				remaining -= tokens
				result.Impressions = append(result.Impressions, i)
			}
		}
	}

	relevantTrimmed := 0

	if topic != "" && scope == "" {
		exclude := make([]string, len(always))

		for i, m := range always {
			exclude[i] = fmt.Sprintf("memory/%d", m.Identifier)
		}

		relevant, f := s.SearchRelevant(topic, 20, exclude)

		if f != nil {
			return nil, nil, fmt.Errorf(
				"%w: %v",
				constant.ErrorRelevantSearch,
				f,
			)
		}

		for _, r := range relevant {
			tokens := s.countTokens(r)

			if remaining-tokens < 0 {
				relevantTrimmed++

				continue
			}

			remaining -= tokens
			result.Relevant = append(result.Relevant, r)
		}
	}

	var d *ProfileDetail

	if detail {
		d = &ProfileDetail{
			Budget:             constant.ProfileBudget,
			AlwaysTokens:       alwaysTokens,
			IndexTokens:        s.countTokens(result.Index),
			IndexTrimmed:       indexTrimmed,
			CompletionTokens:   s.countTokens(result.Completions),
			CompletionsTrimmed: completionsTrimmed,
			ImpressionTokens:   s.countTokens(result.Impressions),
			ImpressionsTrimmed: impressionsTrimmed,
			RelevantTokens:     s.countTokens(result.Relevant),
			RelevantTrimmed:    relevantTrimmed,
			TotalTokens:        constant.ProfileBudget - remaining,
		}
	}

	return result, d, nil
}
