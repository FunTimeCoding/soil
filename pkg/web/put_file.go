package web

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func PutFile(
	locator string,
	header map[string]string,
	b []byte,
) (int, string) {
	var u = new(bytes.Buffer)
	u.Write(b)
	r := NewPutBytes(locator, u)

	for k, v := range header {
		r.Header.Add(k, v)
	}

	p := stall_client.New().Build()
	defer p.Close()
	s := Send(p.Client(), r)
	defer errors.PanicClose(s.Body)

	return s.StatusCode, ReadString(s)
}
