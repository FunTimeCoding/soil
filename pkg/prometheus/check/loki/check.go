package loki

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/check/loki/option"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Check(o *option.Loki) {
	c := loki.NewEnvironment(false)
	f := constant.ColorFormat.Copy()

	if o.Copyable {
		f.Tag(constant.TagCopyable)
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
