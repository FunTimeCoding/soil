package response

import "time"

type Module struct {
	Path     string
	Version  string
	Time     *time.Time
	Indirect bool
	Main     bool
	Update   *Module
}
