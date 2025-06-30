package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type TemplateParam struct {
	Name    string `yaml:"name"`
	Prompt  string `yaml:"prompt,omitempty"`
	Default string `yaml:"default,omitempty"`
}

type TemplateConfig struct {
	Params  []TemplateParam         `yaml:"params"`
	Folders map[string]string       `yaml:"folders,omitempty"`
	Files   map[string]FileTemplate `yaml:"files"`
}

type FileTemplate struct {
	Path     string `yaml:"path"`
	Template string `yaml:"template"`
}

type Config struct {
	Templates map[string]TemplateConfig `yaml:"templates"`
}

// LoadConfig загружает YAML-конфиг из файла
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

// Пути к конфигам
const (
	GlobalConfigDir  = "~/.lekalo"             // Глобальная папка
	GlobalConfigFile = "lekalo_templates.yml"  // Глобальный конфиг
	LocalConfigFile  = ".lekalo_templates.yml" // Локальный конфиг
)

// FindConfigs ищет все конфиги (глобальный + локальный)
func FindConfigs() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	globalPath := filepath.Join(home, ".lekalo", GlobalConfigFile)

	var paths []string
	if _, err := os.Stat(globalPath); err == nil {
		paths = append(paths, globalPath)
	}

	if localPath, err := filepath.Abs(LocalConfigFile); err == nil {
		if _, err := os.Stat(localPath); err == nil {
			paths = append(paths, localPath)
		}
	}

	return paths, nil
}

// MergeConfigs объединяет конфиги (локальные перекрывают глобальные)
func MergeConfigs(configs []*Config) *Config {
	merged := &Config{Templates: make(map[string]TemplateConfig)}
	for _, cfg := range configs {
		for name, tpl := range cfg.Templates {
			merged.Templates[name] = tpl
		}
	}
	return merged
}

func LoadConfigs() (*Config, error) {
	paths, err := FindConfigs()
	if err != nil {
		return nil, err
	}

	var configs []*Config
	for _, path := range paths {
		cfg, err := LoadConfig(path)
		if err != nil {
			return nil, err
		}
		configs = append(configs, cfg)
	}

	return MergeConfigs(configs), nil
}
