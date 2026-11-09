package backup

import "time"

func New(
	instance string,
	name string,
	at time.Time,
	size int64,
) *Backup {
	return &Backup{Instance: instance, Name: name, Time: at, Size: size}
}
