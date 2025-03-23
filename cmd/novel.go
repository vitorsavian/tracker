/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// novelCmd represents the novel command
var novelCmd = &cobra.Command{
	Use:   "book",
	Short: "A command to track my books",
	Long:  `It is not a very complex command to write a long description`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(novelCmd)

	novelCmd.Flags().String("name", "", "name of the novel")
	novelCmd.Flags().Int("page", 0, "page that I'm currently in")
	novelCmd.Flags().Bool("finished", true, "flag for defining if the novel/book is still launching")
}
