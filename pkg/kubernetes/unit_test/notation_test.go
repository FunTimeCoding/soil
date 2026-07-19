package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/notation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestToUnstructured(t *testing.T) {
	assert.Any(
		t,
		&unstructured.Unstructured{Object: map[string]any{"kind": "Pod"}},
		notation.ToUnstructured([]byte(`{"kind":"Pod"}`)),
	)
}
