package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func Download(
	locator string,
	output string,
) {
	p := stall_client.New().Build()
	defer p.Close()
	r, e := p.Client().Get(locator)
	errors.PanicOnError(e)
	defer errors.LogClose(r.Body)
	errors.PanicStatus(r)
	f := system.Create(output)
	defer errors.LogClose(f)
	system.Copy(r.Body, f)
}
