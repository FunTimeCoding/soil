package refusal

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"net/http"
)

func NewDaemonTable() *Table {
	return NewTable().
		With(http.StatusBadRequest, errorsConstant.ValidationClass).
		With(http.StatusUnprocessableEntity, errorsConstant.ValidationClass).
		With(http.StatusForbidden, errorsConstant.ForbiddenClass)
}
