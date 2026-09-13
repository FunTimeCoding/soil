package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"os"
	"path/filepath"
	"testing"
)

func TestScanFileSplitsIdentifiers(t *testing.T) {
	d := t.TempDir()
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(d, "sample.go"),
			[]byte(
				"noiseWithMatch\nsmtpd_banner\nghcr.io/aquasecurity/x\n"+
					"naïveHandler\n",
			),
			0644,
		),
	)
	usage := map[string]*dictionary.WordUsage{
		"match":        dictionary.NewWordUsage("match", "Test", false),
		"banner":       dictionary.NewWordUsage("banner", "Test", false),
		"aquasecurity": dictionary.NewWordUsage("aquasecurity", "Test", false),
		"aquasec":      dictionary.NewWordUsage("aquasec", "Test", false),
		"absent":       dictionary.NewWordUsage("absent", "Test", false),
		"naïve":        dictionary.NewWordUsage("naïve", "Test", false),
	}
	dictionary.ScanFile(filepath.Join(d, "sample.go"), usage)
	assert.True(t, usage["match"].Used)
	assert.True(t, usage["banner"].Used)
	assert.True(t, usage["aquasecurity"].Used)
	assert.True(t, usage["naïve"].Used)
	assert.False(t, usage["aquasec"].Used)
	assert.False(t, usage["absent"].Used)
}

func TestScanFileSplitsDigits(t *testing.T) {
	d := t.TempDir()
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(d, "sample.md"),
			[]byte("aarch64 and microk8s and qcow2\n"),
			0644,
		),
	)
	usage := map[string]*dictionary.WordUsage{
		"aarch":  dictionary.NewWordUsage("aarch", "Test", false),
		"microk": dictionary.NewWordUsage("microk", "Test", false),
		"qcow":   dictionary.NewWordUsage("qcow", "Test", false),
		"absent": dictionary.NewWordUsage("absent", "Test", false),
	}
	dictionary.ScanFile(filepath.Join(d, "sample.md"), usage)
	assert.True(t, usage["aarch"].Used)
	assert.True(t, usage["microk"].Used)
	assert.True(t, usage["qcow"].Used)
	assert.False(t, usage["absent"].Used)
}
