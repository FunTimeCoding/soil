package web

import "github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"

func coverageServer(
	servers []*coverage.Server,
	name string,
) *coverage.Server {
	for _, e := range servers {
		if e.Name == name {
			return e
		}
	}

	return nil
}
