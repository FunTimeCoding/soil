package refusal

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"net/http"
)

func NewTable() *Table {
	return &Table{
		rows: map[int]string{
			http.StatusNotFound:           errorsConstant.NotFoundClass,
			http.StatusConflict:           errorsConstant.ConflictClass,
			http.StatusTooManyRequests:    errorsConstant.UnavailableClass,
			http.StatusBadGateway:         errorsConstant.UnavailableClass,
			http.StatusServiceUnavailable: errorsConstant.UnavailableClass,
			http.StatusGatewayTimeout:     errorsConstant.UnavailableClass,
		},
	}
}
