package gorilla

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func home(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		homeTemplate.Execute(
			w,
			locator.New(
				r.Host,
			).Scheme(constant.Socket).Path(location.Echo).String(),
		),
	)
}
