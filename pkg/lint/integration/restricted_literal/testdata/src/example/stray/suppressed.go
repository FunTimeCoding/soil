package stray

import "example/fakehttp"

func Suppressed() *fakehttp.Client {
	return &fakehttp.Client{Name: "tolerated"} // goanalyze:ignore
}
