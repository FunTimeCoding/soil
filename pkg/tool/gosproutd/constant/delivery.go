package constant

import "time"

type Delivery string

const (
	DeliveryDrop      Delivery = "drop"
	DeliveryQueue     Delivery = "queue"
	DeliveryImmediate Delivery = "immediate"
	DeliveryHold      Delivery = "hold"
)

const (
	QuietWindow = 3 * time.Minute
	WakeWindow  = time.Hour
)
