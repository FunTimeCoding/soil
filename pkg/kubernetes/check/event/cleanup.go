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

	for _, e := range k.EventsSimple(true, false) {
		if deleteOld(k, e, age) {
			continue
		}
	}

	for _, e := range k.EventsSimple(false, true) {
		if deleteIrrelevant(k, e) {
			continue
		}

		if deleteOld(k, e, age) {
			continue
		}
	}
}
