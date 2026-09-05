package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"
)

func (t *Tester) MuteRules() []mute_rule.MuteRule {
	result, e := t.Service.MuteRules()
	errors.PanicOnError(e)

	return result
}
