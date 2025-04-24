package config

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Paths  []string `yaml:"paths"`
	Ignore []string `yaml:"ignore"`
	SortBy string   `yaml:"sort_by"`
	Fetch  bool     `yaml:"fetch_origin"`
	Theme  struct {
		Name   string            `yaml:"name,omitempty"`   // t.ex. "lolcat"
		Colors map[string]string `yaml:"colors,omitempty"` // override
	} `yaml:"theme"`
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(usr.HomeDir, ".gitdash.yaml")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("failed to read config file: " + err.Error())
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.New("invalid config format: " + err.Error())
	}

	// Expand ~ if present
	for i, p := range cfg.Paths {
		cfg.Paths[i] = expandPath(p)
	}
	for i, p := range cfg.Ignore {
		cfg.Ignore[i] = expandPath(p)
	}

	return &cfg, nil
}

func expandPath(p string) string {
	if len(p) > 1 && p[:2] == "~/" {
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, p[2:])
		}
	}
	return p
}
