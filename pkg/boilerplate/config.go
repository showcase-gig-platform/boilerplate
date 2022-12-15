package boilerplate

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Project        string   `yaml:"project"`
	Targets        []Target `yaml:"targets"`
	IgnorePrefixes []string `yaml:"ignore_prefixes"`
}

type Target struct {
	Name   string `yaml:"name"`
	String string `yaml:"string"`
}

var configFiles = []string{
	"boilerplate.yaml",
	"boilerplate.yml",
}

func SearchConfig(dir string) (*Config, error) {
	for _, file := range configFiles {
		path := dir + string(filepath.Separator) + file
		if _, err := os.Stat(path); err == nil {
			return LoadConfig(path)
		}
	}
	return nil, fmt.Errorf("config file not found")
}

func LoadConfig(path string) (*Config, error) {
	conf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var c Config
	if err := yaml.Unmarshal(conf, &c); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &c, nil
}
