package silence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"time"
)

func Check(o *option.Silence) {
	c := common.Alertmanager()
	silences := collect(c)

	if o.Notation {
		printNotation(silences, o)

		return
	}

	if o.Set != "" {
		fmt.Printf("Set: %s\n", c.MustSimpleSilence(o.Set))
	}

	o2 := advanced_option.New()
	o2.All = true
	a, _ := c.MustAlerts(o2, nil)
	fmt.Printf("Alerts: %d\n", len(a))

	if !o.All {
		silences = silence.FilterPermanent(silences)
	}

	var relevant int
	f := constant.Format

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	t := time.Now()

	for _, e := range silences {
		if !o.All && e.State != constant.ActiveState {
			continue
		}

		relevant++
		fmt.Println(e.Format(f))

		if m := matcher.Matches(e, a, t); len(m) > 0 {
			fmt.Printf("  Matching: %d\n", len(m))
		}
	}

	if !o.All && relevant == 0 {
		fmt.Printf(
			"No relevant %s, %d in total\n",
			item.GoSilence.Plural,
			len(silences),
		)
	}
}
