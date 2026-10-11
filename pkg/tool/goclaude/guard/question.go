package guard

import "strings"

func Question(command string) string {
	calls, e := gitCalls(command)

	if e != nil {
		return ""
	}

	var reasons []string

	for _, c := range calls {
		if a := c.Ask(); a != "" {
			reasons = append(reasons, a)
		}
	}

	return strings.Join(reasons, "; ")
}
