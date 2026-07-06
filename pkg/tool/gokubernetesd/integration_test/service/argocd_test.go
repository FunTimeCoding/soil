//go:build local

package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"testing"
)

func TestArgocdApplications(t *testing.T) {
	s := service_tester.New(t)
	s.AddArgoApp("sentry", "sentry", "Synced", "Healthy")
	s.AddArgoApp("bot", "bot", "OutOfSync", "Degraded")
	result, e := s.Service.ArgocdApplications(context.Background(), "test")
	assert.Nil(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "bot", result[0].Name)
	assert.String(t, "OutOfSync", result[0].Sync)
	assert.String(t, "Degraded", result[0].Health)
	assert.String(t, "sentry", result[1].Name)
	assert.String(t, "Synced", result[1].Sync)
}

func TestArgocdApplicationDetail(t *testing.T) {
	s := service_tester.New(t)
	s.AddArgoApp("model-context", "model-context", "Synced", "Healthy")
	result, e := s.Service.ArgocdApplication(
		context.Background(),
		"test",
		"model-context",
		false,
	)
	assert.Nil(t, e)
	assert.NotNil(t, result.Application)
}
