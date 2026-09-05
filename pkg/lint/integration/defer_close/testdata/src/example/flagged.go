package example

import "os"

func Flagged() {
	f, _ := os.Open("test")
	defer f.Close()
}
