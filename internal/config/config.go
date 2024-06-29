package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

var (
	DEBUGKey         = "DEBUG"
	DEBUGShort       = "D"
	ConfigFileKey    = "cfg-file"
	ConfigFileShort  = "c"
	LanguageKey      = "language"
	LanguageShort    = "l"
	DataPathKey      = "data-path"
	DataPathShort    = "d"
	GrammarFileKey   = "grammar-file"
	GrammarFileShort = "g"
	StartPointKey    = "start-point"
	StartPointShort  = "s"
)

type Config struct {
	Language    string `mapstructure:"language"`
	DataPath    string `mapstructure:"data-path"`
	GrammarFile string `mapstructure:"grammar-file"`
	StartPoint  string `mapstructure:"start-point"`
}

var DefaultConfig = Config{
	Language:    "en",
	DataPath:    filepath.Join(xdg.ConfigHome, "randonamer", "data"),
	GrammarFile: "_grammar",
	StartPoint:  "start",
}

func DefaultConfigPath() string {
	return filepath.Join(xdg.ConfigHome, "randonamer")
}

func DefaultConfigFilePath(configPath string) string {
	return filepath.Join(configPath, "config.yaml")
}

func WriteDefaultConfig(configFilePath string) {
	viper.Set(LanguageKey, DefaultConfig.Language)
	viper.Set(DataPathKey, DefaultConfig.DataPath)
	viper.Set(GrammarFileKey, DefaultConfig.GrammarFile)
	viper.Set(StartPointKey, DefaultConfig.StartPoint)

	if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModePerm); err != nil {
		fmt.Printf("Error creating config directory: %v\n", err)
		os.Exit(1)
	}

	if err := viper.WriteConfigAs(configFilePath); err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		os.Exit(1)
	}
}
