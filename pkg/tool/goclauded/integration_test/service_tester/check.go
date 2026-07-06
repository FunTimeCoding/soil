package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/check_result"
)

func (o *Tester) Check(sessionIdentifier string) *check_result.Result {
	result, e := o.Service.Check(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
