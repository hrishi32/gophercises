package cmd

import (
	"fmt"
	"os"
	"strconv"

	"../db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var tasksToDo []int

		for _, arg := range args {
			argInt, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}

			tasksToDo = append(tasksToDo, argInt)

		}

		for _, key := range tasksToDo {
			db.DeleteTask(key)
			fmt.Printf("You have completed the \"%d\" task.", key)
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
