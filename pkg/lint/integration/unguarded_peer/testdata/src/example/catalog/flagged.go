package catalog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Unguarded(title string) error {
	return s.client.Shelve(context.Background(), title)
}

func Bare(
	c Client,
	g *peer.Guard,
) func(context.Context, string) (string, error) {
	return c.Book
}

func (s *Service) shared(
	x context.Context,
	title string,
) error {
	return s.client.Shelve(x, title)
}

func (s *Service) Mixed(title string) error {
	e := s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			return s.shared(x, title)
		},
	)

	if e != nil {
		return e
	}

	return s.shared(context.Background(), title)
}

func (s *Service) Exported(
	x context.Context,
	title string,
) error {
	return s.client.Shelve(x, title)
}

func (s *Service) UseExported(title string) error {
	return s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			return s.Exported(x, title)
		},
	)
}

func fake(
	g *peer.Guard,
	call func() error,
) error {
	return call()
}

func (s *Service) Faked(title string) error {
	return fake(
		s.guard,
		func() error {
			return s.client.Shelve(context.Background(), title)
		},
	)
}

func (s *Service) Retried(title string) error {
	return peer.Retry(
		context.Background(),
		func(x context.Context) error {
			return s.client.Shelve(x, title)
		},
	)
}
