package cmd

import (
	"fmt"
	"strings"

	db "../db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the arguments

		taskToAdd := strings.Join(args, " ")

		_, err := db.CreateTask(taskToAdd)
		if err != nil {
			panic(err)
		}

		fmt.Println("Added", taskToAdd, "to your task list")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
