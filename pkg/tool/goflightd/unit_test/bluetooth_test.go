package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/bluetooth"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"testing"
)

func TestBluetoothParse(t *testing.T) {
	result := bluetooth.Parse(constant.BluetoothSample)
	assert.Integer(t, 2, len(result))
	assert.String(t, "connected", result["Example Mouse"])
	assert.String(t, "disconnected", result["Example Keyboard"])
}

func TestBluetoothParseInvalid(t *testing.T) {
	assert.Integer(t, 0, len(bluetooth.Parse("not json")))
}
