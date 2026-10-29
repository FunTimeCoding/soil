package event

import (
	"github.com/funtimecoding/soil/pkg/kubernetes"
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"time"
)

func cleanup(
	k *client.Client,
	o *option.Event,
	age time.Duration,
) {
	if !o.Clean && !kubernetes.AutoCleanup() {
		return
	}

	for _, v := range k.MustEventsSimple(true, false) {
		if deleteOld(k, v, age) {
			continue
		}
	}

	for _, v := range k.MustEventsSimple(false, true) {
		if deleteIrrelevant(k, v) {
			continue
		}

		if deleteOld(k, v, age) {
			continue
		}
	}
}
