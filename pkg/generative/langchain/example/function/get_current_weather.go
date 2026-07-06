package function

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
)

func getCurrentWeather(
	location string,
	unit string,
) string {
	weather := map[string]any{
		"location":    location,
		"temperature": "6",
		"unit":        unit,
		"forecast":    []string{"sunny", "windy"},
	}

	if unit == "fahrenheit" {
		weather["temperature"] = 43
	}

	b, e := json.Marshal(weather)
	errors.PanicOnError(e)

	return string(b)
}
