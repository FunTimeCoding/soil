package event

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"time"
)

func Print(o *option.Event) {
	k := client.NewEnvironment()
	cleanup(k, o, 7*24*time.Hour)
	events := k.EventsSimple(false, true)

	if o.Notation {
		printNotation(events, o)

		return
	}

	f := constant.Dense.Copy()

	for _, e := range events {
		fmt.Println(e.Format(f))
	}
}
