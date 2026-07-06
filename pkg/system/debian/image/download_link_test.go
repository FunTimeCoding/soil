package image

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestDownloadLink(t *testing.T) {
	assert.String(
		t,
		"https://cdimage.debian.org/cdimage/release/1.2.3/arm64/iso-cd",
		DirectoryLink(semver.New("1.2.3"), constant.ARM64),
	)
}
