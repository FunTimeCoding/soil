package alert

import "time"

type Alert struct {
	Identifier  uint
	Fingerprint string
	Name        string
	Severity    string
	Summary     string
	Labels      map[string]string
	Start       time.Time
	End         *time.Time
}
