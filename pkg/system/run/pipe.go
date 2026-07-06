package run

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
)

func (r *Run) Pipe(
	input string,
	s ...string,
) (string, string) {
	c := r.build(s...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	pipe, e := c.StdinPipe()
	errors.PanicOnError(e)
	errors.PanicOnError(c.Start())
	written, e := pipe.Write([]byte(input))
	errors.PanicOnError(e)
	errors.PanicOnError(pipe.Close())

	if f := c.Wait(); f != nil {
		r.Error = f

		if r.Panic {
			log.Panicf("wait: %e", f)
		}

		return stdout.String(), stderr.String()
	}

	if written == 0 {
		log.Panic("no bytes written")
	}

	return stdout.String(), stderr.String()
}
