package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type Config struct {
	Language    string `mapstructure:"language"`
	DataPath    string `mapstructure:"data_path"`
	GrammarFile string `mapstructure:"grammar_file"`
}

var DefaultConfig = Config{
	Language:    "en",
	DataPath:    filepath.Join(xdg.ConfigHome, "randonamer", "data"),
	GrammarFile: "_grammar",
}

func DefaultConfigPath() string {
	return filepath.Join(xdg.ConfigHome, "randonamer")
}

func DefaultConfigFilePath(configPath string) string {
	return filepath.Join(configPath, "config.yaml")
}

func WriteDefaultConfig(configFilePath string) {
	viper.Set("language", DefaultConfig.Language)
	viper.Set("data_path", DefaultConfig.DataPath)
	viper.Set("grammar_file", DefaultConfig.GrammarFile)

	if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModePerm); err != nil {
		fmt.Printf("Error creating config directory: %v\n", err)
		os.Exit(1)
	}

	if err := viper.WriteConfigAs(configFilePath); err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		os.Exit(1)
	}
}
