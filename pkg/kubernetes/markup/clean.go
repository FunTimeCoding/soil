package markup

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"sigs.k8s.io/yaml"
)

func Clean(v []byte) []byte {
	var result map[string]any
	errors.PanicOnError(yaml.Unmarshal(v, &result))

	return Encode(RemoveEmpty(result).(map[string]any))
}
