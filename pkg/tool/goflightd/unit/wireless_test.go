package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/wireless"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"testing"
)

func TestWirelessParse(t *testing.T) {
	result := wireless.Parse(constant.WirelessSample)
	assert.Integer(t, 6, len(result))
	assert.String(t, "en0 (Wi-Fi)", result["network.Primary IPv4"])
	assert.String(t, "aa:bb:cc:dd:ee:ff", result["wifi.MAC Address"])
	assert.String(t, "-45 dBm", result["wifi.RSSI"])
	assert.String(t, "5g44/80", result["wifi.Channel"])
	assert.String(t, "Yes", result["awdl.AWDL Enabled"])
	assert.String(
		t,
		"{(44, 80MHz), (6, 20MHz)}",
		result["awdl.Channel Sequence"],
	)
}

func TestWirelessParseSkipsPreamble(t *testing.T) {
	result := wireless.Parse("Total Time : ignored\n")
	assert.Integer(t, 0, len(result))
}
