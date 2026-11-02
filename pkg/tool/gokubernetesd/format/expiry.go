package format

import (
	"fmt"
	"time"
)

func Expiry(notAfter string) string {
	if notAfter == "" {
		return ""
	}

	t, e := time.Parse(time.RFC3339, notAfter)

	if e != nil {
		return ""
	}

	days := int(time.Until(t).Hours() / 24)

	if days < 0 {
		return fmt.Sprintf("expired %dd ago", -days)
	}

	return fmt.Sprintf("%dd", days)
}
