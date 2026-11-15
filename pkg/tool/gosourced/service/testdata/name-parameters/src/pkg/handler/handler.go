package handler

import "context"

type Server struct{}

func (s *Server) Get(
	_ context.Context,
	identifier string,
) string {
	return identifier
}

func (s *Server) List(context.Context) []string {
	return nil
}

func Pair(context.Context, string) {}

func Named(y context.Context) context.Context {
	return y
}

func Literal() func(context.Context) string {
	return func(_ context.Context) string {
		return ""
	}
}
