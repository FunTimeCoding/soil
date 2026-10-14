package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func FromFacts(summaries []*fact.Summary) *Summaries {
	result := &Summaries{
		arranges: make(map[string]bool),
		disposes: make(map[string]map[int]bool),
		releases: make(map[string]map[string]bool),
		shared:   make(map[string]bool),
		lends:    make(map[string]bool),
		self:     make(map[string]bool),
		spent:    make(map[string]bool),
		wraps:    make(map[string]map[int]bool),
	}
	delegate := make(map[string][]string)
	forward := make(map[string][]fact.Forward)
	wrapping := make(map[string][]fact.Forward)
	arranging := make(map[string][]fact.Condition)
	spentVia := make(map[string][]string)
	lendsVia := make(map[string][]string)
	release := make(map[string][]fact.Release)

	for _, s := range summaries {
		for _, k := range s.Arranges {
			result.arranges[k] = true
		}

		for k, callees := range s.Delegates {
			delegate[k] = append(delegate[k], callees...)
		}

		for k, arguments := range s.Disposes {
			for _, a := range arguments {
				mark(result.disposes, k, a)
			}
		}

		for k, all := range s.Forwards {
			forward[k] = append(forward[k], all...)
		}

		for k, all := range s.Wraps {
			wrapping[k] = append(wrapping[k], all...)
		}

		for k, all := range s.Arranging {
			arranging[k] = append(arranging[k], all...)
		}

		for _, k := range s.Shared {
			result.shared[k] = true
		}

		for _, k := range s.Lends {
			result.lends[k] = true
		}

		for _, k := range s.Self {
			result.self[k] = true
		}

		for _, k := range s.Spent {
			result.spent[k] = true
		}

		for k, callees := range s.SpentVia {
			spentVia[k] = append(spentVia[k], callees...)
		}

		for k, callees := range s.LendsVia {
			lendsVia[k] = append(lendsVia[k], callees...)
		}

		for k, all := range s.Releases {
			if result.releases[k] == nil {
				result.releases[k] = make(map[string]bool)
			}

			release[k] = append(release[k], all...)
		}
	}

	for changed := true; changed; {
		changed = false

		for f, all := range forward {
			for _, w := range all {
				if result.holdsAll(w.Conditions) &&
					mark(result.disposes, f, w.Argument) {
					changed = true
				}
			}
		}

		for f, all := range wrapping {
			for _, w := range all {
				if result.holdsAll(w.Conditions) &&
					mark(result.wraps, f, w.Argument) {
					changed = true
				}
			}
		}

		for k, all := range release {
			for _, r := range all {
				if result.holds(&r.Condition) &&
					markField(result.releases, k, r.Field) {
					changed = true
				}
			}
		}

		for f, callees := range lendsVia {
			if !result.lends[f] && allKnown(result.lends, callees) {
				result.lends[f] = true
				changed = true
			}
		}

		for f, callees := range spentVia {
			if !result.spent[f] && allKnown(result.spent, callees) {
				result.spent[f] = true
				changed = true
			}
		}
	}

	for f, all := range arranging {
		for _, c := range all {
			if result.holds(&c) {
				result.arranges[f] = true
			}
		}
	}

	for changed := true; changed; {
		changed = false

		for f, all := range delegate {
			if result.arranges[f] {
				continue
			}

			for _, g := range all {
				if result.arranges[g] {
					result.arranges[f] = true
					changed = true

					break
				}
			}
		}
	}

	return result
}
