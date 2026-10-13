package client

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Configuration(
	key string,
	h *http.Client,
) *client.Config {
	return &client.Config{
		ApiKey:         key,
		HttpClient:     h,
		LogLevel:       logrus.ErrorLevel,
		RetryCount:     2,
		RequestTimeout: 10 * time.Second,
	}
}
