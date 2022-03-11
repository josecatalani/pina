package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pina",
	Short: "A cat-friendly invoice generator.",
	Long:  "A cat-friendly invoice generator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pina.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
