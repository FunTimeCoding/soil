package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/budget"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
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

	parentChildren := map[int64][]string{}

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

		parentChildren[parent] = names
	}

	notice := ""
	var completed []client.SearchResult

	if scope == "" {
		results, g := s.ListCompletions()

		if n := s.queryNotice(g); n != "" {
			notice = n
		} else {
			completed = results
		}
	}

	var matched []record.SearchResult

	if topic != "" && scope == "" {
		exclude := make([]string, len(always))

		for i, m := range always {
			exclude[i] = fmt.Sprintf("memory/%d", m.Identifier)
		}

		var g error
		matched, g = s.SearchRelevant(topic, 20, exclude)

		if n := s.queryNotice(g); n != "" {
			notice = n
		}
	}

	b := budget.New(constant.ProfileByteBudget)

	if notice != "" {
		b.Open("")
		b.Charge(notice)
	}

	mark := b.Spent()
	b.Open(profile.Opening(constant.AlwaysSectionHeading))

	for i := range always {
		if children, found := parentChildren[always[i].Identifier]; found {
			always[i].Children = children
		}

		b.Charge(format.AlwaysMemory(&always[i]))
	}

	alwaysBytes := b.Spent() - mark
	mark = b.Spent()
	b.Open(profile.Opening(constant.IndexSectionHeading))
	indexTrimmed := 0
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

		if children, found := parentChildren[m.Identifier]; found {
			m.Children = children
		}

		if !b.Fit(format.IndexEntry(&m)) {
			indexTrimmed++

			continue
		}

		index = append(index, m)
	}

	indexBytes := b.Spent() - mark
	mark = b.Spent()
	b.Open(profile.Opening(constant.CompletionSectionHeading))
	completionsTrimmed := 0
	var completions []*profile.Completion

	for _, r := range completed {
		name := r.Path

		if i := strings.LastIndex(name, stringsConstant.Slash); i >= 0 {
			name = name[:i]
		}

		body := ""

		if r.Body != nil {
			body = *r.Body
		}

		if !b.Fit(format.Completion(name, body)) {
			completionsTrimmed++

			continue
		}

		completions = append(completions, profile.NewCompletion(name, body))
	}

	completionBytes := b.Spent() - mark
	mark = b.Spent()
	b.Open(profile.Opening(constant.ImpressionSectionHeading))
	impressionsTrimmed := 0
	var impressions []record.Impression

	if scope == "" {
		latest, g := s.LatestImpressions(10)

		if g == nil {
			for i := range latest {
				if !b.Fit(format.Impression(&latest[i])) {
					impressionsTrimmed++

					continue
				}

				impressions = append(impressions, latest[i])
			}
		}
	}

	impressionBytes := b.Spent() - mark
	mark = b.Spent()
	b.Open(profile.Opening(constant.RelevantSectionHeading))
	relevantTrimmed := 0
	var relevant []record.SearchResult

	for i := range matched {
		if !b.Fit(format.RelevantMemory(&matched[i])) {
			relevantTrimmed++

			continue
		}

		relevant = append(relevant, matched[i])
	}

	relevantBytes := b.Spent() - mark
	result := profile.NewResult(
		always,
		index,
		relevant,
		impressions,
		completions,
	).WithNotice(notice)

	if !detail {
		return result, nil, nil
	}

	d := profile.NewDetail()
	d.Budget = constant.ProfileByteBudget
	d.AlwaysBytes = alwaysBytes
	d.IndexBytes = indexBytes
	d.IndexTrimmed = indexTrimmed
	d.CompletionBytes = completionBytes
	d.CompletionsTrimmed = completionsTrimmed
	d.ImpressionBytes = impressionBytes
	d.ImpressionsTrimmed = impressionsTrimmed
	d.RelevantBytes = relevantBytes
	d.RelevantTrimmed = relevantTrimmed
	d.TotalBytes = b.Spent()

	return result, d, nil
}
