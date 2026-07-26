package constant

import "time"

const (
	StoreStatusRunning = "running"
	StoreStatusSuccess = "success"
	StoreStatusError   = "error"
	StoreRetentionDays = 14
	StoreRetentionAge  = StoreRetentionDays * 24 * time.Hour
)
