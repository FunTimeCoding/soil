package service_tester

import "context"

func (t *Tester) context() context.Context {
	return context.Background()
}
