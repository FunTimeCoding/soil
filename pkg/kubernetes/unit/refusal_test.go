package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	"testing"
)

func TestAMissingJobIsNotFound(t *testing.T) {
	_, e := refusingClientset(
		kubernetes.NewNotFound(constant.JobResource, "alfa"),
	).Job("default", "alfa")
	assert.True(t, not_found.Is(e))
}

func TestARefusedListIsForbidden(t *testing.T) {
	_, e := refusingClientset(
		kubernetes.NewForbidden(
			constant.JobResource,
			"",
			errors.New("no permission"),
		),
	).Jobs("default")
	assert.True(t, forbidden.Is(e))
}

func TestAConflictingWriteIsConflict(t *testing.T) {
	e := refusingClientset(
		kubernetes.NewConflict(
			constant.JobResource,
			"alfa",
			errors.New("modified"),
		),
	).DeleteJob("default", "alfa")
	assert.True(t, conflict.Is(e))
}

func TestAMissingJobIsNotRunning(t *testing.T) {
	running, e := refusingClientset(
		kubernetes.NewNotFound(constant.JobResource, "alfa"),
	).JobRunning("default", "alfa")
	assert.FatalOnError(t, e)
	assert.False(t, running)
}

func TestAMissingJobIsAlreadyDeleted(t *testing.T) {
	assert.FatalOnError(
		t,
		refusingClientset(
			kubernetes.NewNotFound(constant.JobResource, "alfa"),
		).DeleteJob("default", "alfa"),
	)
}
