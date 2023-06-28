/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var msg string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "To Add, Commit And Push Code To Origin's Main Branch",
	Long: `Use this command to add all files of repo to staging area,
commit them and push them to the origin's main branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		push()
	},
}

func push() {
	commands := []*exec.Cmd{
		exec.Command("git", "add", "."),
		exec.Command("git", "commit", "-m", msg),
		exec.Command("git", "push", "origin", "main"),
	}

	// Execute commands sequentially

	for _, cmd := range commands {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(commitCmd)
	
	commitCmd.Flags().StringVarP(&msg, "msg", "m", "", "commit message")

	err := commitCmd.MarkFlagRequired("msg")
	if err != nil {
		fmt.Println(err)
	}
}
