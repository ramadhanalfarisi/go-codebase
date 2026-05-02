package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "codebase",
	Short: "Go Codebase is a simple and clean codebase for Go",
	Long:  `Go Codebase is a simple and clean codebase for Go with clean architecture and best practices`,
}

func init() {
	RootCmd.AddCommand(ApiCmd)
	RootCmd.AddCommand(MigrateCmd)
	RootCmd.AddCommand(GraphqlCmd)
}
