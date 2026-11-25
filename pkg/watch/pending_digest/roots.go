package pending_digest

import "time"

func (p *Pending) Roots(now time.Time) []string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	seen := map[string]bool{}
	var result []string

	for identifier, root := range p.roots {
		if seen[root] || now.Before(p.next[identifier]) {
			continue
		}

		seen[root] = true
		result = append(result, root)
	}

	return result
}
