package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gosaltd/option"

func (o *Tester) daemonOption() *option.Salt {
	result := option.New()
	result.LitePath = o.litePath
	result.KeySyncInterval = o.keySyncInterval
	result.Master = o.master

	return result
}
