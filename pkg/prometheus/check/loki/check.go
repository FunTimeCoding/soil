package loki

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	checkOption "github.com/funtimecoding/soil/pkg/prometheus/check/loki/option"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Check(o *checkOption.Loki) {
	c := loki.NewEnvironment(false)
	f := option.Color.Copy()

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	if o.Namespace == "" {
		entries := collectOverview(c, o.Namespaces, o.Since)
		printOverview(entries, f)

		return
	}

	messages := collect(
		c,
		o.Namespace,
		o.Since,
		o.Route,
		o.Message,
		o.Exclude,
		o.Limit,
	)

	if o.BodyOnly {
		printBody(messages)

		return
	}

	printLog(messages, f)
}
