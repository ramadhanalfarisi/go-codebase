package cmd

import (
	"github.com/ramadhanalfarisi/go-codebase/app/graphql"
	"github.com/spf13/cobra"
)

var GraphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Start the GraphQL server",
	Long:  `Start the GraphQL server with all the necessary configurations and dependencies`,
	Run: func(cmd *cobra.Command, args []string) {
		api := graphql.NewGraphQL()
		api.Run()
	}}