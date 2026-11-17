package edge

func Down(
	host string,
	reason string,
) *Edge {
	return &Edge{Host: host, Down: true, Reason: reason}
}
