package rerank_benchmark

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"time"
)

func RerankBenchmark() {
	start := time.Now()
	r, e := rerank.New(
		environment.Required(constant.ModelEnvironment),
		environment.Required(constant.TokenizerEnvironment),
	)
	errors.PanicOnError(e)

	defer errors.PanicClose(r)
	fmt.Printf(
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
		fmt.Printf(
			"%s: %v (first score %.4f, last score %.4f)\n",
			c.label,
			time.Since(start),
			results[0].Score,
			results[len(results)-1].Score,
		)
	}
}
