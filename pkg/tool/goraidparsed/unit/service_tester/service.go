package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/service"
	"path/filepath"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		o.service = service.New(
			o.parser,
			filepath.Join(o.t.TempDir(), "TW5_Top_Stat_Parse.html"),
			o.Output,
			logger.New(o.t.Context()),
		)
	}

	return o.service
}
