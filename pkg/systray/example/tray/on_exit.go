//go:build local

package tray

import "fmt"

func onExit() {
	fmt.Println("onExit")
}
