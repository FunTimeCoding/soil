package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/download"
)

func (c *Client) DownloadFile(
	l string,
	path string,
) (*download.Download, error) {
	parameters := map[string]any{"url": l}

	if path != "" {
		parameters["filename"] = path
	}

	r, e := c.send("download_file", parameters)

	if e != nil {
		return download.Stub(), fmt.Errorf("download file: %w", e)
	}

	var result *download.Download

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return download.Stub(), fmt.Errorf("download file: %w", f)
	}

	return result, nil
}
