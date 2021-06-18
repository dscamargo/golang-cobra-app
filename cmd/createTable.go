package cmd

import (
	"github.com/dscamargo/cobraapp/app/scripts"

	"github.com/spf13/cobra"
)

var createTableCmd = &cobra.Command{
	Use:   "createTable",
	Short: "Create a new table in sqlite",
	Long:  `Create a new table in sqlite`,
	Run: func(cmd *cobra.Command, args []string) {
		tableName := cmd.Flag("tableName").Value.String()
		dbName := cmd.Flag("dbName").Value.String()
		scripts.CreateTable(dbName, tableName)
	},
}

func init() {
	rootCmd.AddCommand(createTableCmd)
	createTableCmd.Flags().String("tableName", "user", "Table Name")
	createTableCmd.Flags().String("dbName", "test.db", "Sqlite file name")
}
