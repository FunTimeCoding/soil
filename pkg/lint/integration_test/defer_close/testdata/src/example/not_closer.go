package example

func NotCloser() {
	c := &Client{}
	defer c.Close()
}
