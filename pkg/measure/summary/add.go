package summary

import "github.com/funtimecoding/soil/pkg/measure/count"

func (s *Summary) Add(c *count.Count) {
	s.Files++
	s.Count.Add(c)
}
