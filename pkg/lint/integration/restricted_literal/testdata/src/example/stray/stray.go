package stray

import "example/fakehttp"

func Direct() *fakehttp.Client {
	return &fakehttp.Client{Name: "stray"}
}

func Unrelated() fakehttp.Other {
	return fakehttp.Other{Name: "other"}
}
