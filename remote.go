package git

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

var (
	ErrRemoteUndeclaredURL = errors.New("remote url is empty")
)

type Remote struct {
	c *config.RemoteConfig
}

func NewRemote(c *config.RemoteConfig) *Remote {
	return &Remote{c: c}
}

func (r *Remote) Push(o *PushOptions) error {
	return r.PushContext(context.Background(), o)
}

func (r *Remote) PushContext(ctx context.Context, o *PushOptions) error {
	var url string
	if len(r.c.PushURLs) > 0 {
		url = r.c.PushURLs[0]
	} else if len(r.c.URLs) > 0 {
		url = r.c.URLs[0]
	} else {
		return ErrRemoteUndeclaredURL
	}

	// Mock push transport session setup using the resolved url
	_, err := transport.NewEndpoint(url)
	return err
}

type PushOptions struct {
	Auth transport.AuthMethod
}
