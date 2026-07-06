package alert

import "github.com/funtimecoding/soil/pkg/system"

func (a *Alert) ResolveHosts() []string {
	if a.HostResolved() {
		return []string{a.FullHost()}
	}

	result := system.Lookup(a.FullHost())

	if len(result) == 0 {
		return []string{}
	}

	return result
}
