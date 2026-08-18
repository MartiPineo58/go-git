package git

import (
	"context"
	"testing"

	"github.com/go-git/go-git/v5/config"
	"github.com/stretchr/testify/assert"
)

func TestRemote_Push_PrioritizesPushURL(t *testing.T) {
	cfg := &config.RemoteConfig{
		Name:     "origin",
		URLs:     []string{"https://github.com/example/repo.git"},
		PushURLs: []string{"git@github.com:example/repo.git"},
	}

	r := NewRemote(cfg)
	err := r.PushContext(context.Background(), &PushOptions{})
	assert.NoError(t, err)
}

func TestRemote_Push_FallbackToURL(t *testing.T) {
	cfg := &config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/example/repo.git"},
	}

	r := NewRemote(cfg)
	err := r.PushContext(context.Background(), &PushOptions{})
	assert.NoError(t, err)
}
