package cmd

import (
	"fmt"
	"os"
	"randonamer/internal/config"
	"randonamer/internal/generator"
	"randonamer/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg     config.Config
	cfgFile string
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:   "randonamer",
	Short: "A coolname generator",
	Long: `A coolname generator with support for many languages
and the possibility to use custom configuration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.DebugLog("Starting the generation process")
		coolname, err := generator.GenerateCoolname(cfg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(coolname)
		util.DebugLog("Finished the generation process")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, config.ConfigFileKey, config.ConfigFileShort, "", "path to custom configuration file")
	rootCmd.PersistentFlags().BoolVarP(&debug, config.DEBUGKey, config.DEBUGShort, false, "enable debug logging")
	rootCmd.PersistentFlags().StringP(config.LanguageKey, config.LanguageShort, "", "language to generate coolname")
	rootCmd.PersistentFlags().StringP(config.DataPathKey, config.DataPathShort, "", "path to data directory")
	rootCmd.PersistentFlags().StringP(config.GrammarFileKey, config.GrammarFileShort, "", "path to grammar file")
	rootCmd.PersistentFlags().StringP(config.StartPointKey, config.StartPointShort, "", "start point for grammar")

	viper.BindPFlag(config.ConfigFileKey, rootCmd.PersistentFlags().Lookup(config.ConfigFileKey))
	viper.BindPFlag(config.DEBUGKey, rootCmd.PersistentFlags().Lookup(config.DEBUGKey))
	viper.BindPFlag(config.LanguageKey, rootCmd.PersistentFlags().Lookup(config.LanguageKey))
	viper.BindPFlag(config.DataPathKey, rootCmd.PersistentFlags().Lookup(config.DataPathKey))
	viper.BindPFlag(config.GrammarFileKey, rootCmd.PersistentFlags().Lookup(config.GrammarFileKey))
	viper.BindPFlag(config.StartPointKey, rootCmd.PersistentFlags().Lookup(config.StartPointKey))
}

func initConfig() {
	util.SetDebug(debug)
	configPath := config.DefaultConfigPath()
	configFilePath := config.DefaultConfigFilePath(configPath)
	viper.SetConfigFile(configFilePath)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		if cfgFile == "" {
			config.WriteDefaultConfig(configFilePath)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Can't read config: %v\n", err)
		os.Exit(1)
	}

	viper.SetDefault(config.LanguageKey, config.DefaultConfig.Language)
	viper.SetDefault(config.DataPathKey, config.DefaultConfig.DataPath)
	viper.SetDefault(config.GrammarFileKey, config.DefaultConfig.GrammarFile)
	viper.SetDefault(config.StartPointKey, config.DefaultConfig.StartPoint)

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Can't unmarshal config: %v\n", err)
		os.Exit(1)
	}

	expandConfigPaths()
	overrideConfigWithFlags()
	util.DebugLog("Configuration initialized: %+v", cfg)
}

func expandConfigPaths() {
	cfg.DataPath = os.ExpandEnv(cfg.DataPath)
	cfg.GrammarFile = os.ExpandEnv(cfg.GrammarFile)
}

func overrideConfigWithFlags() {
	if viper.IsSet(config.LanguageKey) && viper.GetString(config.LanguageKey) != "" {
		cfg.Language = viper.GetString(config.LanguageKey)
	}
	if viper.IsSet(config.DataPathKey) && viper.GetString(config.DataPathKey) != "" {
		cfg.DataPath = os.ExpandEnv(viper.GetString(config.DataPathKey))
	}
	if viper.IsSet(config.GrammarFileKey) && viper.GetString(config.GrammarFileKey) != "" {
		cfg.GrammarFile = os.ExpandEnv(viper.GetString(config.GrammarFileKey))
	}
	if viper.IsSet(config.StartPointKey) && viper.GetString(config.StartPointKey) != "" {
		cfg.StartPoint = viper.GetString(config.StartPointKey)
	}
}
