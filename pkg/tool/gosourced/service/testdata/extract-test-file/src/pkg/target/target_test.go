//go:build ci

package target

import (
	"example/pkg/go-widget"
	"testing"
)

func newWidget(name string) *widget.Widget {
	return &widget.Widget{Name: name}
}

func TestWidgetName(t *testing.T) {
	if newWidget("alfa").Name != "alfa" {
		t.Error("name did not survive construction")
	}
}
