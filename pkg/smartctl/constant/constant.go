package constant

import "regexp"

const (
	ScanCommand   = "sudo -n smartctl --scan -j"
	ReportCommand = "sudo -n smartctl -j -a %s"
)

const UnreadMask = 0b11

var DevicePath = regexp.MustCompile(`^/dev/[a-z0-9]+$`)
