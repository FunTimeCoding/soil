package defaulted

type Client struct {
	clients map[string]*Client
}

func Stub() *Client {
	return &Client{clients: make(map[string]*Client)}
}
