package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func FromFacts(summaries []*fact.Summary) *Summaries {
	result := &Summaries{arranges: make(map[string]bool)}
	delegate := make(map[string][]string)

	for _, s := range summaries {
		for _, k := range s.Arranges {
			result.arranges[k] = true
		}

		for k, callees := range s.Delegates {
			delegate[k] = append(delegate[k], callees...)
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
