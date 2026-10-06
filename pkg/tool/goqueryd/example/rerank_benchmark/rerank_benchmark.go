package rerank_benchmark

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func RerankBenchmark() {
	root := environment.Required(constant.RerankDirectoryEnvironment)

	for _, m := range constant.RerankModels {
		if !system.DirectoryExists(join.Join(root, m.Name)) {
			console.Format("%s: absent from %s, skipped\n", m.Name, root)

			continue
		}

		window := environment.FallbackInteger(
			constant.RerankSequenceEnvironment,
			m.SequenceLength,
		)

		if window > m.MaximumLength {
			console.Format(
				"%s: window %d beyond its maximum %d, skipped\n",
				m.Name,
				window,
				m.MaximumLength,
			)

			continue
		}

		benchmarkModel(m)
	}
}
