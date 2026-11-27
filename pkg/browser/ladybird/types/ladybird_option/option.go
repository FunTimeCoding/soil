package ladybird_option

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"

type Option struct {
	DriverPath string
	Address    string
	Profile    string
	Headless   bool
	Event      func(*message.Reply)
}
