package fetch

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
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
			t.Label = t.Identifier

			for _, prefix := range []string{alert, silence, event, file} {
				if strings.HasPrefix(t.Identifier, prefix) {
					t.Label = fmt.Sprintf("%s%d", prefix, i+1)

					break
				}
			}
		}

		return result
	}
}
