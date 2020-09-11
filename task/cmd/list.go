package cmd

import (
	"fmt"

	db "../db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListAllTasks()

		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to do!")
			return
		}
		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			fmt.Println(task.Key, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
