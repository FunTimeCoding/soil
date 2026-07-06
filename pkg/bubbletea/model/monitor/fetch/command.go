package fetch

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"strings"
)

func Command() tea.Cmd {
	return func() tea.Msg {
		result := Message{}

		for _, c := range List() {
			if i := Run(c); len(i) > 0 {
				result.Items = append(result.Items, i...)
			}
		}

		alert := fmt.Sprintf("%s-", constant.GoAlert.Prefix)
		silence := fmt.Sprintf("%s-", constant.GoSilence.Prefix)
		event := fmt.Sprintf("%s-", constant.GoKevt.Prefix)
		file := fmt.Sprintf("%s-", constant.GoFile.Prefix)

		for i, t := range result.Items {
			if strings.HasPrefix(t.Identifier, alert) {
				t.Identifier = fmt.Sprintf("%s%d", alert, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, silence) {
				t.Identifier = fmt.Sprintf("%s%d", silence, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, event) {
				t.Identifier = fmt.Sprintf("%s%d", event, i+1)

				continue
			}

			if strings.HasPrefix(t.Identifier, file) {
				t.Identifier = fmt.Sprintf("%s%d", file, i+1)

				continue
			}
		}

		return result
	}
}
