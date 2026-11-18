package reading

func NewDetached(
	server string,
	battery string,
) *Reading {
	return &Reading{Server: server, Battery: battery}
}
