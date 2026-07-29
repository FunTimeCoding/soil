package gorilla

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func home(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		constant.GorillaHomeTemplate.Execute(
			w,
			locator.New(
				r.Host,
			).Scheme(webConstant.Socket).Path(webConstant.LocationEcho).String(),
		),
	)
}
