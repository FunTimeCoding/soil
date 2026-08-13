package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
)

func main() {
	if false {
		var White = console.NewColor("#fff")
		fmt.Println(White("W"))
	}

	if false {
		for r := range 15 {
			for g := range 15 {
				for b := range 15 {
					fmt.Printf(
						"%s",
						console.NewColor(fmt.Sprintf("#%x%x%x", r, g, b))("W"),
					)
				}

				println()
			}
		}
	}

	if true {
		c := console.Gradient("#00ff00", "#ff0000", 250)

		for i := range len(c) {
			fmt.Printf("%s", c[i]("W"))
		}

		println()
	}
}
