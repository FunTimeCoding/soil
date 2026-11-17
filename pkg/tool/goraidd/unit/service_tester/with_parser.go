package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed/client/mock_client"

func WithParser(p *mock_client.Client) Option {
	return func(o *Tester) {
		o.Parser = p
	}
}
