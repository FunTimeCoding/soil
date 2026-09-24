package summary

import "github.com/funtimecoding/soil/pkg/measure/count"

func New(languageName string) *Summary {
	return &Summary{Language: languageName, Count: count.New()}
}
