package cmd

import (
	"github.com/dscamargo/cobraapp/api"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run Http",
	Long:  `Run Http server`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		api.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringP("port", "p", "8080", "Http port")
}
