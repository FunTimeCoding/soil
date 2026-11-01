package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
)

func New(
	from *revision.Revision,
	to *revision.Revision,
) *Report {
	before := byKey(from.Symbols)
	after := byKey(to.Symbols)
	pairs := samePairs(before, after)
	pairs = append(pairs, namePairs(before, after)...)
	pairs = append(pairs, receiverPairs(before, after, pairs)...)
	pairs = append(pairs, packagePairs(before, after, pairs)...)
	pairs = append(pairs, shapePairs(before, after, moves(pairs))...)
	mapping := moves(pairs)
	var result []*entry.Entry

	for _, p := range pairs {
		p.Class = classify(p, mapping)
		result = append(result, p)
	}

	result = append(result, leftovers(before, after)...)
	sortEntries(result)

	return &Report{
		Entries:      result,
		FromUnparsed: from.Unparsed,
		ToUnparsed:   to.Unparsed,
	}
}
