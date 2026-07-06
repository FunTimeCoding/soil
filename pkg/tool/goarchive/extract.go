package goarchive

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/nwaples/rardecode/v2"
	"io"
	"path/filepath"
)

func extract(
	path string,
	output string,
	password string,
) {
	var options []rardecode.Option

	if password != "" {
		options = append(options, rardecode.Password(password))
	}

	r, e := rardecode.OpenReader(path, options...)
	errors.PanicOnError(e)
	defer errors.PanicClose(r)
	count := 0

	for {
		header, e := r.Next()

		if e == io.EOF {
			break
		}

		errors.PanicOnError(e)
		target := filepath.Join(output, header.Name)

		if header.IsDir {
			system.MakeDirectory(target)

			continue
		}

		system.MakeDirectory(filepath.Dir(target))
		f := system.Create(target)
		_, e = io.Copy(f, r)
		errors.PanicOnError(e)
		errors.PanicClose(f)
		count++
	}

	fmt.Printf("%s: %d files\n", output, count)
}
