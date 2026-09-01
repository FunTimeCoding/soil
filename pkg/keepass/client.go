package keepass

import (
	"github.com/tobischo/gokeepasslib/v3"
	"time"
)

type Client struct {
	database *gokeepasslib.Database
	path     string
	loaded   time.Time
}
