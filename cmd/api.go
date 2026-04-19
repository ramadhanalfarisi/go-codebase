package cmd

import (
	"github.com/ramadhanalfarisi/go-codebase/app/api"
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Long:  `Start the API server with all the necessary configurations and dependencies`,
	Run:   func(cmd *cobra.Command, args []string) {
		api := api.NewApi()
		api.Run()
	}}
