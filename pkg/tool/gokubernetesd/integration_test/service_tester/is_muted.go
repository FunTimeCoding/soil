package service_tester

import "github.com/funtimecoding/soil/pkg/errors"

func (t *Tester) IsMuted(
	reason string,
	message string,
	cluster string,
) bool {
	result, e := t.Service.IsMuted(reason, message, cluster)
	errors.PanicOnError(e)

	return result
}
