package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

// 子命令注册
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
