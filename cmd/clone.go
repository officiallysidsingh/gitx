/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var gitUrl string

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "To Clone A Repository And Start NeoVim In It",
	Long: `Use this command to clone a Repository, CD into that Repository
and then open that Repository in Neovim.

**Note: 
  Please Only Enter The SSH Link Of The Repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		clone()
	},
}

func clone() {
	// Getting the repository name
	parts := strings.Split(gitUrl, "/")
	repoName := strings.TrimSuffix(parts[1], ".git")

	// Execute commands sequentially

	// Clone Command
	cmd := exec.Command("git", "clone", gitUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing clone command: %v\n", err)
		return
	}

	// CD Command
	err = os.Chdir(repoName)
	if err != nil {
		fmt.Printf("Error executing cd command: %v\n", err)
		return
	}

	// Nvim Command
	cmd = exec.Command("nvim")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error executing nvim command: %v\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringVarP(&gitUrl, "url", "u", "", "url of the repository")
	err := cloneCmd.MarkFlagRequired("url")
	if err != nil {
		fmt.Println(err)
	}
}
