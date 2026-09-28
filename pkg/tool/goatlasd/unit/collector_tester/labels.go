package collector_tester

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

func Labels(pairs ...string) *[]*client.Label {
	var v []*client.Label

	for i := 0; i < len(pairs); i += 2 {
		v = append(v, &client.Label{Key: pairs[i], Value: pairs[i+1]})
	}

	return &v
}
