package markup

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"sigs.k8s.io/yaml"
)

func Encode(a any) []byte {
	result, e := yaml.Marshal(a)
	errors.PanicOnError(e)

	return result
}
