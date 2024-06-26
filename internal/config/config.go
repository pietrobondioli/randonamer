package config

import (
	"fmt"
	"os"
	"path/filepath"
	"randonamer/internal/util"

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
	DataPath:    filepath.Join(xdg.DataHome, "randonamer", "data"),
	GrammarFile: "_grammar",
}

func InitConfig(cfgFile string, cfg *Config) {
	configPath := filepath.Join(xdg.ConfigHome, "randonamer")
	configFile := "config"
	configFileExt := ".yaml"

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)

	viper.SetDefault("language", DefaultConfig.Language)
	viper.SetDefault("data_path", DefaultConfig.DataPath)
	viper.SetDefault("grammar_file", DefaultConfig.GrammarFile)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(filepath.Join(configPath, configFile+configFileExt))
		viper.SafeWriteConfigAs(filepath.Join(configPath, configFile+configFileExt))
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("Can't unmarshal config:", err)
		os.Exit(1)
	}

	// if err := CopyDefaultConfigFiles(DefaultConfig.DataPath, cfg.DataPath); err != nil {
	// 	fmt.Println("Error copying default config files:", err)
	// 	os.Exit(1)
	// }
}

func CopyDefaultConfigFiles(srcDir, dstDir string) error {
	return util.CopyDir(srcDir, dstDir)
}
