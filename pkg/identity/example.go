package identity

func Example() *Tool {
	return New("example", "", "example").WithStamp("", "", "")
}
