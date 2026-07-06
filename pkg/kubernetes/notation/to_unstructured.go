package notation

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ToUnstructured(b []byte) *unstructured.Unstructured {
	result := &unstructured.Unstructured{}
	errors.PanicOnError(result.UnmarshalJSON(b))

	return result
}
