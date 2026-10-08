package response

func NewSubscription(
	root string,
	label string,
	channel string,
	lastEvent string,
) *Subscription {
	return &Subscription{
		Root:      root,
		Label:     label,
		Channel:   channel,
		LastEvent: lastEvent,
	}
}
