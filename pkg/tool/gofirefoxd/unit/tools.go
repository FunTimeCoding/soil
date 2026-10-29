package unit

import (
	"github.com/funtimecoding/soil/pkg/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/tool_case"
)

func tools() []tool_case.Tool {
	tab := map[string]any{"tab_id": 1}

	return []tool_case.Tool{
		{
			Name:      constant.ListTabs,
			Arguments: map[string]any{},
			Method:    func(c *mock_client.Client) any { return c.Tabs },
		},
		{
			Name:      constant.ReadTab,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.ReadTab },
		},
		{
			Name:      constant.CreateTab,
			Arguments: map[string]any{"url": "https://example.com"},
			Method:    func(c *mock_client.Client) any { return c.CreateTab },
		},
		{
			Name:      constant.CloseTab,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.CloseTab },
		},
		{
			Name:      constant.Navigate,
			Arguments: map[string]any{"tab_id": 1, "url": "https://example.com"},
			Method:    func(c *mock_client.Client) any { return c.Navigate },
		},
		{
			Name:      constant.NavigateBack,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.NavigateBack },
		},
		{
			Name:      constant.GroupTabs,
			Arguments: map[string]any{"tab_ids": []int{1}},
			Method:    func(c *mock_client.Client) any { return c.GroupTabs },
		},
		{
			Name:      constant.UngroupTab,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.UngroupTab },
		},
		{
			Name:      constant.UpdateGroup,
			Arguments: map[string]any{"group_id": 1},
			Method:    func(c *mock_client.Client) any { return c.UpdateGroup },
		},
		{
			Name:      constant.WaitFor,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.WaitFor },
		},
		{
			Name:      constant.Evaluate,
			Arguments: map[string]any{"tab_id": 1, "expression": "1"},
			Method:    func(c *mock_client.Client) any { return c.Evaluate },
		},
		{
			Name:      constant.Click,
			Arguments: map[string]any{"tab_id": 1, "selector": "#go"},
			Method:    func(c *mock_client.Client) any { return c.Click },
		},
		{
			Name:      constant.Fill,
			Arguments: map[string]any{"tab_id": 1, "selector": "#q"},
			Method:    func(c *mock_client.Client) any { return c.Fill },
		},
		{
			Name:      constant.Screenshot,
			Arguments: tab,
			Method:    func(c *mock_client.Client) any { return c.Screenshot },
		},
		{
			Name:      constant.DownloadFile,
			Arguments: map[string]any{"url": "https://example.com/a"},
			Method:    func(c *mock_client.Client) any { return c.DownloadFile },
		},
		{
			Name:      constant.DownloadStatus,
			Arguments: map[string]any{},
			Method:    func(c *mock_client.Client) any { return c.DownloadStatus },
		},
	}
}
