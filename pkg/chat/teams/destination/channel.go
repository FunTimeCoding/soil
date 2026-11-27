package destination

func Channel(
	team string,
	channel string,
) *Destination {
	return &Destination{Team: team, Channel: channel}
}
