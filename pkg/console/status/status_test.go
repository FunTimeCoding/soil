package status_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/status/extended"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/console/status/tagged"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"testing"
)

func TestStatus(t *testing.T) {
	apple := extended.New(1, "a", "b")
	assert.String(
		t,
		"1 | a | b\n  line1\n  line2",
		apple.Format(option.New().Extended()),
	)
}

func TestStatusNested(t *testing.T) {
	apple := extended.New(1, "a", "b")
	orange := extended.New(2, "c", "d")
	f := option.New().Extended()
	appleOutput := fmt.Sprintf("%s\n", apple.Format(f))
	assert.String(t, "1 | a | b\n  line1\n  line2\n", appleOutput)
	orangeOutput := fmt.Sprintf("%s\n", orange.Format(f.Indent(1)))
	assert.String(
		t,
		"  2 | c | d\n    line1\n    line2\n",
		orangeOutput,
	)

	if false {
		assert.String(
			t,
			"1 | a | b\n  line1\n  line2\n  2 | c | d\n    line1\n    line2\n",
			key_value.Empty(appleOutput, orangeOutput),
		)
	}
}

func TestTagLine(t *testing.T) {
	apple := tagged.New(1, "a", "b")
	assert.String(
		t,
		"1 | a\n  line1\n  line2\n  RawList: &{String:b}",
		apple.Format(option.New().Extended().Tag(tag.Usage).Raw()),
	)
}
