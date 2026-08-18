package config

import (
	"github.com/go-git/go-git/v5/config/format"
)

type Decoder struct {
	d *format.Decoder
}

func NewDecoder(d *format.Decoder) *Decoder {
	return &Decoder{d: d}
}

func (d *Decoder) Decode(cfg *Config) error {
	// Minimal decoder implementation for remote parsing
	for _, s := range d.d.Sections {
		if s.Name == "remote" {
			if err := d.decodeRemote(s, cfg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Decoder) decodeRemote(s *format.Section, cfg *Config) error {
	name := s.Subsection
	r, ok := cfg.Remotes[name]
	if !ok {
		r = &RemoteConfig{Name: name}
		cfg.Remotes[name] = r
	}

	for _, o := range s.Options {
		switch o.Key {
		case "url":
			r.URLs = append(r.URLs, o.Value)
		case "pushurl":
			r.PushURLs = append(r.PushURLs, o.Value)
		case "fetch":
			r.Fetch = append(r.Fetch, RefSpec(o.Value))
		}
	}

	return nil
}
