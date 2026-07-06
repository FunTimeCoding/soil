package helper

import "github.com/funtimecoding/soil/pkg/web/locator"

func Base(host string) string {
	return locator.New(host).String()
}
