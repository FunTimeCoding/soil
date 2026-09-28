package outpost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	outpostConstant "github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewOptional(l *logger.Logger) *Collector {
	hosts := environment.Optional(constant.OutpostHostsEnvironment)
	token := environment.Optional(outpostConstant.TokenEnvironment)

	if hosts == "" || token == "" {
		return nil
	}

	port := environment.RequiredInteger(outpostConstant.PortEnvironment)
	insecure := environment.Exists(outpostConstant.InsecureEnvironment)
	var targets []*Target

	for _, host := range split.Comma(hosts) {
		u := locator.New(host)

		if port != 0 {
			u.Port(port)
		}

		if insecure {
			u.Insecure()
		}

		c, e := client.NewClientWithResponses(
			u.String(),
			client.WithRequestEditorFn(web.BearerEditor(token)),
		)
		errors.PanicOnError(e)
		targets = append(targets, NewTarget(host, c))
	}

	return New(targets, l)
}
