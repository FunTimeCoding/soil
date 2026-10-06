package rerank_benchmark

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_model"
	"time"
)

func benchmarkModel(m *rerank_model.Model) {
	start := time.Now()
	r, e := rerank.NewModel(m)
	errors.PanicOnError(e)
	console.Format("== %s, sequence %d\n", m.Name, r.SequenceLength())

	defer errors.PanicClose(r)
	console.Format(
		"New (tokenizer + runtime + session): %v\n",
		time.Since(start),
	)

	for _, c := range []benchmarkCase{
		{label: "batch 30 distinct", documents: distinctDocuments(30)},
		{label: "batch 30 distinct", documents: distinctDocuments(30)},
		{label: "batch 7 distinct", documents: distinctDocuments(7)},
		{label: "batch 1", documents: distinctDocuments(1)},
		{
			label:     "batch 30 identical (deduplicated)",
			documents: identicalDocuments(30),
		},
	} {
		start = time.Now()
		results, f := r.Rank("database migration strategy", c.documents)
		errors.PanicOnError(f)
		console.Format(
			"%s: %v (first score %.4f, last score %.4f)\n",
			c.label,
			time.Since(start),
			results[0].Score,
			results[len(results)-1].Score,
		)
	}
}
