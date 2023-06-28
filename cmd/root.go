/*
Copyright © 2023 SIDDHARTH SINGH siddharthsingh2014@gmail.com
*/package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitx",
	Short: "An Easy Way To Perform Git Operations",
	Long: `This CLI App helps us perform a collection of git commands by running just a single command.
Running just a single command will add, commit and push the code into origin's main branch.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
