/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// novelCmd represents the novel command
var novelCmd = &cobra.Command{
	Use:   "novel",
	Short: "A command to track my books",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var novelCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var novelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var novelUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var novelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(novelCmd)

	novelCmd.AddCommand(novelCreateCmd)
	novelCmd.AddCommand(novelDeleteCmd)
	novelCmd.AddCommand(novelUpdateCmd)
	novelCmd.AddCommand(novelGetCmd)

	novelCreateCmd.Flags().String("name", "", "name of the novel")
	novelCreateCmd.Flags().Int("page", 0, "page that I'm currently in")
	novelCreateCmd.Flags().Bool("finished", true, "flag for defining if the novel/book is still launching")
}
