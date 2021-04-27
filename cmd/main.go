package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/NeoReids/backend-tryonline-golang/app"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "tryonline",
	Short: "tryonline backend command",
	Run: func(cmd *cobra.Command, args []string) {
		app.InitApp()
	},
}

var migrationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "for migrate all schema",
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
}

var seedCommand = &cobra.Command{
	Use: "seed",
	Short: "Seed data from default seeder.",
	Run: func(cmd *cobra.Command, args []string) {
		Seed()
	},
}

func Execute() {
	rootCommand.AddCommand(migrationCommand)
	rootCommand.AddCommand(seedCommand)
	if errorCommand := rootCommand.Execute(); errorCommand != nil {
		fmt.Fprintln(os.Stderr, errorCommand)
		os.Exit(1)
	}
}
