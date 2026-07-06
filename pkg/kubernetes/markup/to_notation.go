package markup

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func ToNotation(filePath string) []byte {
	f := system.Open(filePath)
	defer errors.LogClose(f)
	var a any
	errors.PanicOnError(yaml.NewYAMLToJSONDecoder(f).Decode(&a))

	return notation.Marshal(a)
}
