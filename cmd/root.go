package cmd

import (
	"fmt"
	"os"
	"randonamer/internal/config"
	"randonamer/internal/generator"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg     config.Config
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "randonamer",
	Short: "A coolname generator",
	Long: `A coolname generator with support for many languages
and the possibility to use custom configuration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		coolname, err := generator.GenerateCoolname(cfg)
		if err != nil {
			fmt.Println(err)
		}

		println(coolname)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "cfgFile", "c", "", "path to custom configuration file")
	rootCmd.PersistentFlags().StringP("language", "l", "", "language to generate coolname")
	rootCmd.PersistentFlags().String("data_path", "", "path to data directory")
	rootCmd.PersistentFlags().String("grammar_file", "", "path to grammar file")

	viper.BindPFlag("language", rootCmd.PersistentFlags().Lookup("language"))
	viper.BindPFlag("data_path", rootCmd.PersistentFlags().Lookup("data_path"))
	viper.BindPFlag("grammar_file", rootCmd.PersistentFlags().Lookup("grammar_file"))
}

func initConfig() {
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

	viper.SetDefault("language", config.DefaultConfig.Language)
	viper.SetDefault("data_path", config.DefaultConfig.DataPath)
	viper.SetDefault("grammar_file", config.DefaultConfig.GrammarFile)

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Can't unmarshal config: %v\n", err)
		os.Exit(1)
	}

	overrideConfigWithFlags()
	fmt.Printf("%+v\n", cfg)
}

func overrideConfigWithFlags() {
	if viper.IsSet("language") && viper.GetString("language") != "" {
		cfg.Language = viper.GetString("language")
	}
	if viper.IsSet("data_path") && viper.GetString("data_path") != "" {
		cfg.DataPath = viper.GetString("data_path")
	}
	if viper.IsSet("grammar_file") && viper.GetString("grammar_file") != "" {
		cfg.GrammarFile = viper.GetString("grammar_file")
	}
}
