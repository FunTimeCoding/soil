package subscription

import "time"

func New(
	callsign string,
	root string,
	channel string,
	alias string,
) *Subscription {
	return &Subscription{
		Callsign:          callsign,
		RootIdentifier:    root,
		ChannelIdentifier: channel,
		Alias:             alias,
		LastEvent:         time.Now(),
	}
}
