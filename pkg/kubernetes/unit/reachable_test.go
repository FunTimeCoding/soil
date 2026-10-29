package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestAnAnsweringClusterIsReachable(t *testing.T) {
	assert.FatalOnError(
		t,
		client.FromInterface(fake.NewClientset(), constant.InCluster).
			Reachable(t.Context()),
	)
}

func TestARefusedVersionIsForbidden(t *testing.T) {
	e := refusingClientset(
		kubernetes.NewForbidden(
			schema.GroupResource{Resource: "version"},
			"",
			errors.New("no permission"),
		),
	).Reachable(t.Context())
	assert.True(t, forbidden.Is(e))
}
