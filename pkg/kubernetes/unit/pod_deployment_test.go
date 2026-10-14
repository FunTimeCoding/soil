package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes"
	"testing"
)

func TestPodDeployment(t *testing.T) {
	assert.String(
		t,
		"runner",
		kubernetes.PodDeployment("runner-6d4b8c9f7-abcde"),
	)
	assert.String(
		t,
		"service-alfa",
		kubernetes.PodDeployment("service-alfa-7b6c5d4e3-klmno"),
	)
}

func TestPodDeploymentTooFewParts(t *testing.T) {
	assert.String(t, "", kubernetes.PodDeployment("work-station"))
	assert.String(t, "", kubernetes.PodDeployment("host"))
	assert.String(t, "", kubernetes.PodDeployment(""))
}
