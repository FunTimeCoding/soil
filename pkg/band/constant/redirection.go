package constant

import "regexp"

var (
	RedirectionInstance = regexp.MustCompile(
		`(?s)<(\w+:)?AMT_RedirectionService\b.*?</(\w+:)?AMT_RedirectionService>`,
	)
	RedirectionOpening = regexp.MustCompile(`^<(\w+):AMT_RedirectionService>`)
	ListenerDisabled   = regexp.MustCompile(
		`(<(?:\w+:)?ListenerEnabled>)false(</(?:\w+:)?ListenerEnabled>)`,
	)
)
