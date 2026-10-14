package constant

import "time"

type Delivery string

const (
	DeliveryDrop Delivery = "drop"
	DeliveryHold Delivery = "hold"
	DeliverySend Delivery = "send"
)

const (
	QuietWindow = 3 * time.Minute
	MaximumHold = 30 * time.Minute
)
