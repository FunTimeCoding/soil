package session_page

import "time"

func promoted(at *time.Time) string {
	if at == nil {
		return ""
	}

	return at.Local().Format("15:04")
}
