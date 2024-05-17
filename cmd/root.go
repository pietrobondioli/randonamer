package cmd

import (
	"fmt"
	"randonamer/internal/generator"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "randonamer",
	Short: "A coolname generator",
	Long: `A coolname generator with support to many languages
  and possibility to use custom configuration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		coolname, err := generator.GenerateCoolname(generator.Config{
			Language:   generator.Language(cmd.Flag("language").Value.String()),
			ConfigPath: cmd.Flag("config").Value.String(),
		})
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
	rootCmd.PersistentFlags().StringP("language", "l", "en", "language to generate coolname")
	rootCmd.PersistentFlags().StringP("config", "c", "", "path to custom configuration file")
}
