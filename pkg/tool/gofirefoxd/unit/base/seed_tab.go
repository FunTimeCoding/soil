package base

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (s *Server) SeedTab() {
	b := tab.New()
	b.Identifier = 100
	b.Locator = "https://example.com"
	b.Title = "Example"
	s.Source.AddTab(b)
}
