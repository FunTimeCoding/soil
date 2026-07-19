package service

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester/pod"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/ambiguous_pods"
	"testing"
)

func TestLogsDirectPod(t *testing.T) {
	s := service_tester.New(t)
	s.AddPod(pod.New("default", "nginx-abc123", "Running"))
	_, e := s.Service.Logs(
		context.Background(),
		"test",
		service.LogsQuery{
			Name:      "nginx-abc123",
			Namespace: "default",
		},
	)
	assert.Nil(t, e)
}

func TestLogsAmbiguousPods(t *testing.T) {
	s := service_tester.New(t)
	labels := map[string]string{"app": "nginx"}
	s.AddDeployment("default", "nginx", 2, 2)
	s.AddPod(pod.New("default", "nginx-abc", "Running").WithLabels(labels))
	s.AddPod(pod.New("default", "nginx-def", "Running").WithLabels(labels))
	_, e := s.Service.Logs(
		context.Background(),
		"test",
		service.LogsQuery{
			Name:      "deployment/nginx",
			Namespace: "default",
		},
	)
	assert.NotNil(t, e)
	var ap *ambiguous_pods.AmbiguousPods
	assert.True(t, errors.As(e, &ap))
	assert.Count(t, 2, ap.Pods)
}

func TestLogsSinglePodDeployment(t *testing.T) {
	s := service_tester.New(t)
	labels := map[string]string{"app": "redis"}
	s.AddDeployment("default", "redis", 1, 1)
	s.AddPod(pod.New("default", "redis-xyz", "Running").WithLabels(labels))
	_, e := s.Service.Logs(
		context.Background(),
		"test",
		service.LogsQuery{
			Name:      "deployment/redis",
			Namespace: "default",
		},
	)
	assert.Nil(t, e)
}
