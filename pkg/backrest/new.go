package backrest

func New(command ...string) *Client {
	return &Client{command: command}
}
