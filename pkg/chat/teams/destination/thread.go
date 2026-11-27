package destination

func Thread(
	team string,
	channel string,
	message string,
) *Destination {
	return &Destination{Team: team, Channel: channel, Thread: message}
}
