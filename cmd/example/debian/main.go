package main

import "github.com/funtimecoding/soil/pkg/system/debian/example"

func main() {
	example.Download()

	if false {
		example.Netboot()
		example.Packer()
	}
}
