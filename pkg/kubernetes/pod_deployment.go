package kubernetes

import "strings"

func PodDeployment(pod string) string {
	parts := strings.Split(pod, "-")

	if len(parts) < 3 {
		return ""
	}

	return strings.Join(parts[:len(parts)-2], "-")
}
