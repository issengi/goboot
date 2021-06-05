package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/issengi/goboot/app"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "goboot",
	Short: "goboot command helper",
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

var makeRepository = &cobra.Command{
	Use:   "repository",
	Short: "command for make repository",
	Run: func(cmd *cobra.Command, args []string) {
		nameRepository := args[0]
		listDirectory := map[string]string{
			fmt.Sprintf(`%s/%s`, nameRepository, `repository`): fmt.Sprintf(`%s_repository.go`,
				nameRepository),
			fmt.Sprintf(`%s/%s`, nameRepository, `usecase`): fmt.Sprintf(`%s_usecase.go`,
				nameRepository),
		}
		for key, item := range listDirectory {
			err := os.MkdirAll(key, 0755)
			if err!=nil{
				panic(err)
			}
			_, err = os.Create(fmt.Sprintf(`%s/%s`, key, item))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

var makeMigrationFile = &cobra.Command{
	Use: "migration",
	Short: "Command for make migration file",
	Run: func(cmd *cobra.Command, args []string) {
		var sqlExtension = `.sql`
		var countFile = 0
		err := filepath.Walk("migrations",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if filepath.Ext(path) == sqlExtension {
					countFile += 1
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
		var prefixNameFile = fmt.Sprintf(`%06d_%s`, countFile/2 + 1, args[0])
		var listFile = []string{
			fmt.Sprintf(`%s/%s.up.sql`, `migrations`, prefixNameFile),
			fmt.Sprintf(`%s/%s.down.sql`, `migrations`, prefixNameFile),
		}
		for _, v := range listFile {
			_, err = os.Create(v)
			if err!=nil{
				panic(err)
			}
		}
	},
}

func newCommandMake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make",
		Short: "still develop",
		Long:  "testing long description",
	}
	return cmd
}

func Execute() {
	makeCommand := newCommandMake()
	makeCommand.AddCommand(makeRepository)
	makeCommand.AddCommand(makeMigrationFile)
	rootCommand.AddCommand(migrationCommand)
	rootCommand.AddCommand(seedCommand)
	if errorCommand := rootCommand.Execute(); errorCommand != nil {
		fmt.Fprintln(os.Stderr, errorCommand)
		os.Exit(1)
	}
}
