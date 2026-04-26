package cmd

import (
	"github.com/ramadhanalfarisi/go-codebase/app/graphql"
	"github.com/spf13/cobra"
)

var GraphqlCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Long:  `Start the API server with all the necessary configurations and dependencies`,
	Run: func(cmd *cobra.Command, args []string) {
		api := graphql.NewGraphQL()
		api.Run()
	}}