package goyaml

import (
	"errors"
	soil "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"go.yaml.in/yaml/v3"
	"io"
)

func Check(path string) error {
	f := system.Open(path)
	defer soil.PanicClose(f)
	d := yaml.NewDecoder(f)

	for {
		var v any
		e := d.Decode(&v)

		if errors.Is(e, io.EOF) {
			return nil
		}

		if e != nil {
			return e
		}
	}
}
