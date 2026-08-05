//go:build linux || darwin

// Package machinery narrates the obvious.
package machinery

import _ "embed"

//go:embed data.txt
var Data string

// Helper returns true.
func Helper() bool {
	// the training could not resist
	return true
}
