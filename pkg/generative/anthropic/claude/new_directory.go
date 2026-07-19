package claude

func NewDirectory(base string) *Client {
	return &Client{base: base}
}
