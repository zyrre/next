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

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.UndoTask(args[0])
		fmt.Println(chalk.Bold.TextStyle(utils.TaskColor + chalk.Bold.TextStyle(args[0]) + utils.TextColor + " undone"))
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
