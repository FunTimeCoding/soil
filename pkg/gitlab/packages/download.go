package packages

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/request"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func Download(
	link string,
	token string,
	outputFile string,
) {
	o := stall_client.New().Build()
	defer o.Close()
	r := web.NewGet(link)
	request.PrivateToken(r, token)
	p := web.Send(o.Client(), r)
	defer errors.PanicClose(p.Body)
	errors.PanicStatus(p)
	f := system.Create(outputFile)
	defer errors.PanicClose(f)
	system.Copy(p.Body, f)
}
