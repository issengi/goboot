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
		if err := Migrate(); err!=nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCommand.AddCommand(migrationCommand)
	if errorCommand := rootCommand.Execute(); errorCommand != nil {
		fmt.Fprintln(os.Stderr, errorCommand)
		os.Exit(1)
	}
}
