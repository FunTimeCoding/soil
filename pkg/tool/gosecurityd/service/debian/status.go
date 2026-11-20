package debian

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
	"slices"
	"time"
)

func (s *Service) Status(
	name string,
	release string,
) (*status_result.Result, error) {
	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	var loaded []*tracker.Tracker

	for _, source := range s.sources {
		if t := source.Tracker(); t != nil {
			loaded = append(loaded, t)
		}
	}

	if len(loaded) == 0 {
		return nil, unavailable.Format(
			"%s not loaded yet",
			vulnerabilityConstant.DebianSubject,
		)
	}

	var releases []string
	listed := release == ""
	var modified time.Time
	var result []*status.Status
	known := false

	for _, t := range loaded {
		releases = append(releases, t.Releases()...)
		listed = listed || t.Lists(release)
		statuses, f := t.Vulnerability(name)

		if f != nil {
			continue
		}

		known = true

		if modified.IsZero() || t.Modified().Before(modified) {
			modified = t.Modified()
		}

		for _, v := range statuses {
			if release == "" || v.Release == release {
				result = append(result, v)
			}
		}
	}

	if !listed {
		slices.Sort(releases)

		return nil, not_found.Format(
			"release not found: %s (tracker lists %s)",
			release,
			join.CommaSpace(releases),
		)
	}

	if !known {
		return nil, not_found.New(vulnerabilityConstant.Kind, name)
	}

	slices.SortFunc(
		result,
		func(a *status.Status, b *status.Status) int {
			return cmp.Or(
				cmp.Compare(a.Package, b.Package),
				cmp.Compare(a.Release, b.Release),
			)
		},
	)

	return status_result.New(name, result, modified), nil
}
