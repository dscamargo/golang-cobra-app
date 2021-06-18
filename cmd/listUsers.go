package cmd

import (
	"github.com/dscamargo/cobraapp/app/scripts"

	"github.com/spf13/cobra"
)

var listUsersCmd = &cobra.Command{
	Use:   "listUsers",
	Short: "List all DB users",
	Long:  `List all users ordered by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		scripts.ListUsers()
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)
}
