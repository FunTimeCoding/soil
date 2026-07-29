package function

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func getCurrentWeather(
	location string,
	unit string,
) string {
	weather := map[string]any{
		"location":                       location,
		constant.LangchainTemperatureKey: "6",
		"unit":                           unit,
		"forecast":                       []string{"sunny", "windy"},
	}

	if unit == "fahrenheit" {
		weather[constant.LangchainTemperatureKey] = 43
	}

	b, e := json.Marshal(weather)
	errors.PanicOnError(e)

	return string(b)
}
