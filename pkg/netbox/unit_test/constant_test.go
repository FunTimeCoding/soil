package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-Hook-Signature", constant.SignatureHeader)
}
