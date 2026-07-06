package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/score"

func (c *Client) Score(
	_ string,
	_ string,
) (*score.Score, error) {
	result := score.Stub()
	result.HP = c.stats.HP
	result.MP = c.stats.MP
	result.XP = c.stats.XP + 10
	result.GP = c.stats.GP + 1
	result.Level = c.stats.Level

	return result, nil
}
