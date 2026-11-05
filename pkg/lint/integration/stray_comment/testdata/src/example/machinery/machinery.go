//go:build linux || darwin

// Package machinery narrates the obvious.
package machinery

import _ "embed"

//go:embed data.txt
var Data string

//goanalyze:by-value measured - copied on every step
type Step struct {
	Value int
}

// Helper returns true.
func Helper() bool {
	// noinspection SpellCheckingInspection
	// renovate: datasource=docker
	// the training could not resist
	return true
}
