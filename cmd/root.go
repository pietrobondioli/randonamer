package cmd

import (
	"fmt"
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
	Long: `A coolname generator with support to many languages
  and possibility to use custom configuration files.`,
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
	viper.BindPFlag("cfgFile", rootCmd.PersistentFlags().Lookup("cfgFile"))

	rootCmd.PersistentFlags().StringP("language", "l", "en", "language to generate coolname")
	viper.BindPFlag("language", rootCmd.PersistentFlags().Lookup("language"))

	rootCmd.PersistentFlags().String("data_path", "", "path to data directory")
	viper.BindPFlag("data_path", rootCmd.PersistentFlags().Lookup("data_path"))

	rootCmd.PersistentFlags().String("grammar_file", "", "path to grammar file")
	viper.BindPFlag("grammar_file", rootCmd.PersistentFlags().Lookup("grammar_file"))
}

func initConfig() {
	config.InitConfig(cfgFile, &cfg)
	fmt.Printf("%+v\n", cfg)
}
