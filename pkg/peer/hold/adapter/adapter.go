package adapter

import "context"

type Adapter struct {
	open  func(context.Context) error
	serve func(context.Context) error
	close func()
}
