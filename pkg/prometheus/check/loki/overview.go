package loki

import "time"

type Overview struct {
	Namespace string
	Count     int
	Latest    time.Time
}
