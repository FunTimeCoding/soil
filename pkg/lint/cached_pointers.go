package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_dependency"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
	"io"
	"strings"
)

func cachedPointers(
	repo *repository.Repository,
	o *option.Lint,
	r *output.Results,
	base *resolver.Resolver,
) Checker {
	return func(
		p string,
		reader io.Reader,
	) *file_report.Report {
		b, e := io.ReadAll(reader)
		errors.PanicOnError(e)
		content := string(b)
		cached, found := o.Resolved.Find(repo.Root, p, content)

		if found && !o.Resolved.Auditing() {
			for _, u := range cached.Unchecked {
				r.AddUnchecked(u.Path, u.Line, u.Span, u.Reason)
			}

			return cached.Report.Clone()
		}

		d := pointer_dependency.New()
		var spans []*unchecked.Unchecked
		result := Pointers(
			recordedResolver(base, d, repo.Siblings),
			func(
				path string,
				line int,
				span string,
				reason constant.Reason,
			) {
				spans = append(spans, unchecked.New(path, line, span, reason))
				r.AddUnchecked(path, line, span, reason)
			},
		)(p, strings.NewReader(content))

		if found {
			before := pointerDigest(cached.Report, cached.Unchecked)

			if after := pointerDigest(result, spans); before != after {
				o.Resolved.Mismatch(repo.Root, p, before, after)
			}
		}

		o.Resolved.Store(repo.Root, p, content, result, spans, d)

		return result
	}
}
