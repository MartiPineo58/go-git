package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
)

const (
	DefaultRefSpec = "+refs/heads/*:refs/remotes/origin/*"
)

var (
	ErrRemoteConfigEmptyName = errors.New("remote config: empty name")
	ErrRemoteConfigEmptyURL  = errors.New("remote config: empty URL")
)

type RemoteConfig struct {
	Name     string    `json:"name"`
	URLs     []string  `json:"urls"`
	PushURLs []string  `json:"pushurls,omitempty"`
	Fetch    []RefSpec `json:"fetch"`
}

func (c *RemoteConfig) Validate() error {
	if c.Name == "" {
		return ErrRemoteConfigEmptyName
	}

	if len(c.URLs) == 0 {
		return ErrRemoteConfigEmptyURL
	}

	for _, r := range c.Fetch {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}
