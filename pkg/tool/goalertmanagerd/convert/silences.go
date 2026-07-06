package convert

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"

func Silences(v []*silence.Silence) []*SlimSilence {
	var result []*SlimSilence

	for _, s := range v {
		result = append(result, Silence(s))
	}

	return result
}
