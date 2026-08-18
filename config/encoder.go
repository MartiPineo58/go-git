package config

import (
	"github.com/go-git/go-git/v5/config/format"
)

type Encoder struct {
	e *format.Encoder
}

func NewEncoder(e *format.Encoder) *Encoder {
	return &Encoder{e: e}
}

func (e *Encoder) Encode(cfg *Config) error {
	// Minimal encoder implementation for remote serialization
	for _, r := range cfg.Remotes {
		s := format.NewSection("remote", r.Name)
		if err := e.encodeRemote(s, r); err != nil {
			return err
		}
		e.e.Sections = append(e.e.Sections, s)
	}
	return nil
}

func (e *Encoder) encodeRemote(s *format.Section, r *RemoteConfig) error {
	for _, url := range r.URLs {
		s.AddOption("url", url)
	}
	for _, url := range r.PushURLs {
		s.AddOption("pushurl", url)
	}
	for _, fetch := range r.Fetch {
		s.AddOption("fetch", string(fetch))
	}
	return nil
}
