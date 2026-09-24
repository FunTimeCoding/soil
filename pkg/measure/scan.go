package measure

import (
	"github.com/funtimecoding/soil/pkg/measure/classifier"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/file"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"github.com/funtimecoding/soil/pkg/measure/result"
	"github.com/funtimecoding/soil/pkg/source"
	"os"
)

func Scan(
	o *option.Measure,
	r *registry.Registry,
) *result.Result {
	out := result.New()
	skips := append(append([]string{}, constant.DefaultSkips...), o.Skips...)

	for _, root := range o.Paths {
		for _, path := range walk(root, skips) {
			b, e := os.ReadFile(path)

			if e != nil || isBinary(b) {
				out.AddSkipped(path)

				continue
			}

			content := string(b)

			if source.IsGeneratedHeader(content) {
				continue
			}

			l := r.ByPath(path)

			if l == nil {
				l = r.ByShebang(firstLine(content))
			}

			if l == nil {
				out.AddUnplaced(path)

				continue
			}

			out.Add(file.New(path, l.Name, classifier.New(l).Classify(content)))
		}
	}

	return out
}
