package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: ``,
	Short: ``,
	Long: ``,
}


func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}