package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/unit/refusing_clientset"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	"testing"
)

func TestMissingJobIsNotFound(t *testing.T) {
	_, e := refusing_clientset.New(
		kubernetes.NewNotFound(constant.JobResource, "alfa"),
	).Job("default", "alfa")
	assert.True(t, not_found.Is(e))
}

func TestRefusedListIsForbidden(t *testing.T) {
	_, e := refusing_clientset.New(
		kubernetes.NewForbidden(
			constant.JobResource,
			"",
			errors.New("no permission"),
		),
	).Jobs("default")
	assert.True(t, forbidden.Is(e))
}

func TestConflictingWriteIsConflict(t *testing.T) {
	e := refusing_clientset.New(
		kubernetes.NewConflict(
			constant.JobResource,
			"alfa",
			errors.New("modified"),
		),
	).DeleteJob("default", "alfa")
	assert.True(t, conflict.Is(e))
}

func TestMissingJobIsNotRunning(t *testing.T) {
	running, e := refusing_clientset.New(
		kubernetes.NewNotFound(constant.JobResource, "alfa"),
	).JobRunning("default", "alfa")
	assert.FatalOnError(t, e)
	assert.False(t, running)
}

func TestMissingJobIsAlreadyDeleted(t *testing.T) {
	assert.FatalOnError(
		t,
		refusing_clientset.New(
			kubernetes.NewNotFound(constant.JobResource, "alfa"),
		).DeleteJob("default", "alfa"),
	)
}
