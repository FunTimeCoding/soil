package console

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
)

func NewColor(hex string) face.SprintFunction {
	return func(
		format string,
		a ...any,
	) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(hex)).Render(
			fmt.Sprintf(
				format,
				a...,
			),
		)
	}
}
