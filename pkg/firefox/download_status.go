package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/download"
)

func (c *Client) DownloadStatus(identifier int) ([]*download.Download, error) {
	parameters := map[string]any{}

	if identifier > 0 {
		parameters["download_id"] = identifier
	}

	r, e := c.send("download_status", parameters)

	if e != nil {
		return nil, fmt.Errorf("download status: %w", e)
	}

	var result []*download.Download

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return nil, fmt.Errorf("download status: %w", f)
	}

	return result, nil
}
