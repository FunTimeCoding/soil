package main

import "github.com/funtimecoding/soil/pkg/system/linux/debian/install/example"

func main() {
	example.Download()

	if false {
		example.Netboot()
		example.Packer()
	}
}
