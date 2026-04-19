package cmd

import (
	"github.com/ramadhanalfarisi/go-codebase/app/migrate"
	"github.com/spf13/cobra"
)


var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations to create or update the database schema`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Migrate()
	}}
