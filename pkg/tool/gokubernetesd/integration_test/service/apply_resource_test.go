package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"testing"
)

var testManifest = `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config
data:
  key: value`

func TestApplyResource(t *testing.T) {
	s := service_tester.New(t)
	result, e := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
		},
	)
	assert.Nil(t, e)
	assert.String(t, "ConfigMap", result.Kind)
	assert.String(t, "test-config", result.Name)
	assert.String(t, "default", result.Namespace)
}

func TestApplyResourceAlreadyExists(t *testing.T) {
	s := service_tester.New(t)
	_, f := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
		},
	)
	assert.Nil(t, f)
	_, e := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
		},
	)
	assert.NotNil(t, e)
}

func TestApplyResourceOverride(t *testing.T) {
	s := service_tester.New(t)
	_, f := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
		},
	)
	assert.Nil(t, f)
	result, e := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
			Override:  true,
		},
	)
	assert.Nil(t, e)
	assert.String(t, "test-config", result.Name)
}

func TestApplyResourceDryRun(t *testing.T) {
	s := service_tester.New(t)
	result, e := s.Service.ApplyResource(
		context.Background(),
		"test",
		service.ApplyQuery{
			Manifest:  testManifest,
			Namespace: "default",
			DryRun:    true,
		},
	)
	assert.Nil(t, e)
	assert.String(t, "test-config", result.Name)
}
