package connection

func New(
	host string,
	token string,
	s Subscriber,
) *Connection {
	return &Connection{
		host:        host,
		token:       token,
		subscriber:  s,
		subscribers: make(map[uint64]Subscriber),
	}
}
