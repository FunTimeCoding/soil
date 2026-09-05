package service_tester

import "github.com/funtimecoding/soil/pkg/errors"

func (t *Tester) Mute(
	reason string,
	message string,
	cluster string,
) {
	errors.PanicOnError(t.Service.MuteEvent(reason, message, cluster))
}
