package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestRolloutRestart(t *testing.T) {
	s := service_tester.New(t)
	s.AddDeployment("default", "nginx", 1, 1)
	e := s.Service.RolloutRestart(
		context.Background(),
		"test",
		"deployment",
		"nginx",
		"default",
	)
	assert.Nil(t, e)
	deployment, f := s.Clientset.AppsV1().Deployments("default").Get(
		context.Background(),
		"nginx",
		v1.GetOptions{},
	)
	assert.Nil(t, f)
	annotations := deployment.Spec.Template.Annotations
	_, hasRestart := annotations["kubectl.kubernetes.io/restartedAt"]
	assert.True(t, hasRestart)
}

func TestRolloutRestartUnsupportedType(t *testing.T) {
	s := service_tester.New(t)
	e := s.Service.RolloutRestart(
		context.Background(),
		"test",
		"service",
		"nginx",
		"default",
	)
	assert.NotNil(t, e)
}
