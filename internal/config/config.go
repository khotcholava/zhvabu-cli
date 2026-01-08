package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Defaults Defaults `json:"defaults"`
	Project  Project  `json:"project"`
}

type Defaults struct {
	Component ComponentDefaults `json:"component"`
}

type ComponentDefaults struct {
	Style          string `json:"style"`
	Path           string `json:"path"`
	SkipStyle      bool   `json:"skipStyle"`
	ComponentStyle string `json:"componentStyle"`
	Memo           bool   `json:"memo"`
	ForwardRef     bool   `json:"forwardRef"`
	Class          bool   `json:"class"`
}

type Project struct {
	Prefix string `json:"prefix"`
}

func ReadConfig() (*Config, error) {
	configFile := "react-cli.json"

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return GetDefaultConfig(), nil
		}

		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func WriteConfig(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("react-cli.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultConfig() *Config {
	return &Config{
		Defaults: Defaults{
			Component: ComponentDefaults{
				Style:          "css",
				Path:           ".",
				SkipStyle:      false,
				ComponentStyle: "functional",
				Memo:           false,
				ForwardRef:     false,
				Class:          false,
			},
		},
		Project: Project{
			Prefix: "",
		},
	}
}
