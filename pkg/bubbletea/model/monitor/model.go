package monitor

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/soil/pkg/bubbletea/types/modal"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"time"
)

type Model struct {
	width          int
	height         int
	auto           bool
	topBar         string
	table          *table.Model
	items          []*item.Item
	bottomBar      string
	second         int
	connect        bool
	monitor        *client.Client
	updates        chan claim.Message
	claims         map[string][]string
	claimError     error
	owner          string
	notice         string
	user           string
	hostname       string
	toast          []*toast.Toast
	nextToast      int
	initialResized bool
	modal          *modal.Modal
	lastFetch      time.Time
}
