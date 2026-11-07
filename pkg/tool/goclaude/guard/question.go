package guard

import "strings"

func Question(command string) string {
	var reasons []string

	if calls, e := gitCalls(command); e == nil {
		for _, c := range calls {
			if a := c.Ask(); a != "" {
				reasons = append(reasons, a)
			}
		}
	}

	if r := ripgrepQuestion(command); r != "" {
		reasons = append(reasons, r)
	}

	return strings.Join(reasons, "; ")
}
