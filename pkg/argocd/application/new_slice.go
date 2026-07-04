package application

import "github.com/funtimecoding/go-library/pkg/argocd/response"

func NewSlice(v []response.Application) []*Application {
	var result []*Application

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
