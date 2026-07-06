package generic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
)

func sortLabels(s []string) []string {
	var result []string

	if slices.Contains(s, constant.NamespaceLabel) {
		result = append(result, constant.NamespaceLabel)
	}

	if slices.Contains(s, constant.PodLabel) {
		result = append(result, constant.PodLabel)
	}

	if slices.Contains(s, constant.ContainerLabel) {
		result = append(result, constant.ContainerLabel)
	}

	for _, e := range s {
		if e == constant.NamespaceLabel ||
			e == constant.PodLabel ||
			e == constant.ContainerLabel {
			continue
		}

		result = append(result, e)
	}

	return result
}
