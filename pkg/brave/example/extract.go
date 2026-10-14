package example

import "github.com/funtimecoding/soil/pkg/generative/openai/site"

func Extract() {
	s := site.New()
	defer s.Close()
	s.ExtractFlow(true)
}
