package packages

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/request"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
)

func Download(
	link string,
	token string,
	outputFile string,
) {
	r := web.NewGet(link)
	request.PrivateToken(r, token)
	p := web.Send(web.Client(), r)
	defer errors.PanicClose(p.Body)
	errors.PanicStatus(p)
	f := system.Create(outputFile)
	defer errors.PanicClose(f)
	system.Copy(p.Body, f)
}
