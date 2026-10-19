package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
)

func UnknownHost() {
	c := opnsense.NewEnvironment()
	defer c.Close()
	identifier := "00000000-0000-0000-0000-000000000000"
	h := request.New()
	h.Host = "alfa"
	h.Domain = "example.test"
	h.Address = "192.0.2.10"
	e := c.SetHost(identifier, h)
	console.Format("set: class=%q error=%v\n", classify.Class(e), e)
	f := c.DeleteHost(identifier)
	console.Format("delete: class=%q error=%v\n", classify.Class(f), f)
}
