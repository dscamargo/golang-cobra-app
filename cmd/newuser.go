package cmd

import (
	"github.com/dscamargo/cobraapp/app/scripts"
	"github.com/spf13/cobra"
)

var newuserCmd = &cobra.Command{
	Use:   "newUser",
	Short: "Create a new user",
	Long:  `Add a new user in users list`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		email := cmd.Flag("email").Value.String()
		scripts.NewUser(name, email)
	},
}

func init() {
	rootCmd.AddCommand(newuserCmd)
	newuserCmd.Flags().String("name", "user", "User name")
	newuserCmd.Flags().String("email", "user@email.com.br", "User email")
}
