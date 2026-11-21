package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"sync"
	"time"
)

type Client struct {
	Closed          int
	failure         error
	vulnerabilities map[string]*client.Vulnerability
	statuses        map[string]*client.DebianStatus
	flags           []*client.Flag
	snapshots       map[string]client.SystemSnapshot
	updated         map[string]time.Time
	details         map[string]*client.SystemDetail
	groups          []*client.FindingGroup
	upgrades        []client.UpgradeGroup
	remediation     []client.RemediationItem
	affected        map[string]*client.AffectedSystems
	scoreStatus     client.ScoreStatus
	mutex           sync.Mutex
}
