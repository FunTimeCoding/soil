package unit

import "time"

type silenceWindowCase struct {
	name          string
	start         time.Time
	end           time.Time
	expectedMatch bool
	description   string
}
