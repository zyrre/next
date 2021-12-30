/*
Copyright © 2021 Peter Krantz zyrree@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"github.com/zyrre/next/utils"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileMap, _ := utils.FileToMap("next.md")
		for _, completed := range fileMap["Completed"] {
			fmt.Println(chalk.Strikethrough.TextStyle(completed))
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
