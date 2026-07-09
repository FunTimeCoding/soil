package image

import (
	"github.com/google/go-github/v89/github"
	"time"
)

type Image struct {
	Identifier int64
	Digest     string
	Tags       []string
	Create     time.Time
	Raw        *github.PackageVersion
}
