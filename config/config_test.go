package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteConfig_PushURL(t *testing.T) {
	cfg := &RemoteConfig{
		Name:     "origin",
		URLs:     []string{"https://github.com/example/repo.git"},
		PushURLs: []string{"git@github.com:example/repo.git"},
	}

	assert.NoError(t, cfg.Validate())
	assert.Equal(t, "git@github.com:example/repo.git", cfg.PushURLs[0])
}
