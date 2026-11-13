package catalog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Client interface {
	Reachable(x context.Context) error
	Close() error
	Address() string
	Book(x context.Context, title string) (string, error)
	Shelve(x context.Context, title string) error
}

type Service struct {
	client Client
	guard  *peer.Guard
}

func guarded[T any](
	g *peer.Guard,
	call func() (T, error),
) (T, error) {
	var result T
	e := g.Attempt(
		context.Background(),
		func(context.Context) error {
			var f error
			result, f = call()

			return f
		},
	)

	return result, e
}

func (s *Service) attempt(call func() error) error {
	return s.guard.Do(
		context.Background(),
		func(context.Context) error {
			return call()
		},
	)
}

func (s *Service) Book(title string) (string, error) {
	return guarded(
		s.guard,
		func() (string, error) {
			return s.client.Book(context.Background(), title)
		},
	)
}

func (s *Service) Shelve(title string) error {
	return s.attempt(
		func() error {
			return s.client.Shelve(context.Background(), title)
		},
	)
}

func (s *Service) Direct(title string) error {
	return s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			return s.client.Shelve(x, title)
		},
	)
}

func (s *Service) lookup(
	x context.Context,
	title string,
) (string, error) {
	return s.client.Book(x, title)
}

func (s *Service) Find(title string) error {
	return s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			_, e := s.lookup(x, title)

			return e
		},
	)
}

func (s *Service) Where() string {
	return s.client.Address()
}

func (s *Service) Stop() error {
	return s.client.Close()
}

func (s *Service) Probe(x context.Context) error {
	return s.client.Reachable(x)
}

func Lookup(
	c Client,
	g *peer.Guard,
) func(context.Context, string) (string, error) {
	return wrap(g, c.Book)
}

func wrap(
	g *peer.Guard,
	call func(context.Context, string) (string, error),
) func(context.Context, string) (string, error) {
	return func(x context.Context, title string) (string, error) {
		var result string
		e := g.Attempt(
			x,
			func(y context.Context) error {
				var f error
				result, f = call(y, title)

				return f
			},
		)

		return result, e
	}
}

type Mock struct{}

func (m *Mock) Reachable(context.Context) error {
	return nil
}

func (m *Mock) Seed(title string) error {
	return nil
}

type Tester struct {
	mock  *Mock
	guard *peer.Guard
}

func (o *Tester) Fill() error {
	return o.mock.Seed("catalog")
}

type Plain struct {
	client Client
}

func (p *Plain) Fetch(title string) (string, error) {
	return p.client.Book(context.Background(), title)
}

func (s *Service) Through(title string) (string, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(x context.Context) (string, error) {
			return s.client.Book(x, title)
		},
	)
}

func (s *Service) Typed(title string) (string, error) {
	return peer.Guarded[string](
		context.Background(),
		s.guard,
		func(x context.Context) (string, error) {
			return s.client.Book(x, title)
		},
	)
}

func (s *Service) via(
	call func(context.Context) (string, error),
) (string, error) {
	return peer.Guarded(context.Background(), s.guard, call)
}

func (s *Service) Via(title string) (string, error) {
	return s.via(
		func(x context.Context) (string, error) {
			return s.client.Book(x, title)
		},
	)
}
