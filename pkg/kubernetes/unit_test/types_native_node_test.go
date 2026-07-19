package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNode(t *testing.T) {
	assert.NotNil(
		t,
		node.New(
			&core.Node{ObjectMeta: meta.ObjectMeta{Name: upper.Alfa}},
			"",
		),
	)
}
