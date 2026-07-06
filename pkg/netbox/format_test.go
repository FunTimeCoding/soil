package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	fixture := []string{"a", "b"}
	assert.String(t, "[a b]", fmt.Sprintf("%v", fixture))
	assert.String(t, "[a b]", fmt.Sprintf("%+v", fixture))
	assert.String(
		t,
		`[]string{"a", "b"}`,
		fmt.Sprintf("%#v", fixture),
	)
}
