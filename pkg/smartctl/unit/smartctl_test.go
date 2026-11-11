package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/smartctl"
	"testing"
)

func TestReportReadsHealthyDrive(t *testing.T) {
	r, e := smartctl.Report(fixture.Read("disk", "smartctl_nvme_healthy.json"))
	assert.FatalOnError(t, e)
	assert.String(t, "/dev/nvme1", r.Device.Name)
	assert.String(t, "Example NVMe 2TB", r.ModelName)
	assert.True(t, r.Status.Passed)
	assert.Integer(t, 49, r.Temperature.Current)
	assert.Integer(t, 0, r.Health.CriticalWarning)
	assert.Integer(t, 100, r.Health.AvailableSpare)
	assert.Integer(t, 0, r.Health.PercentageUsed)
	assert.Integer(t, int64(0), r.Health.MediaErrors)
}

func TestReportKeepsFailingDriveDespiteExitStatus(t *testing.T) {
	r, e := smartctl.Report(
		fixture.Read("disk", "smartctl_nvme_failing_derived.json"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 8, r.Run.ExitStatus)
	assert.False(t, r.Status.Passed)
	assert.Integer(t, 4, r.Health.CriticalWarning)
	assert.Integer(t, int64(3), r.Health.MediaErrors)
}

func TestReportRefusesUnreadDevice(t *testing.T) {
	_, e := smartctl.Report(
		`{"device":{"name":"/dev/nvme9"},"smartctl":{"exit_status":2}}`,
	)
	assert.True(t, unexpected.Is(e))
}

func TestReportRefusesOutputWithoutReport(t *testing.T) {
	_, e := smartctl.Report(`{"json_format_version":[1,0]}`)
	assert.True(t, unexpected.Is(e))
}
