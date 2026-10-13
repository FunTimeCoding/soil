package blessed

import "example/fakehttp"

func New() *fakehttp.Client {
	return &fakehttp.Client{Name: "allowed"}
}
