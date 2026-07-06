package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
)

var White = console.NewColor("#fff")

func main() {
	if false {
		fmt.Println(White("W"))
	}

	if false {
		for r := 0; r < 15; r++ {
			for g := 0; g < 15; g++ {
				for b := 0; b < 15; b++ {
					fmt.Printf(
						"%s",
						console.NewColor(
							fmt.Sprintf(
								"#%x%x%x",
								r,
								g,
								b,
							),
						)("W"),
					)
				}

				println()
			}
		}
	}

	if true {
		c := console.Gradient("#00ff00", "#ff0000", 250)
		count := len(c)

		for i := 0; i < count; i++ {
			fmt.Printf(
				"%s",
				c[i]("W"),
			)
		}

		println()
	}
}
