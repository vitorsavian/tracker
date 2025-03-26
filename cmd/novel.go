/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/controller"
	"github.com/vitorsavian/tracker/pkg/infra/env"
)

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
		env.SetEnv()

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			logrus.Errorf("Unable to get name from cli: %v\n", err)
			return
		}

		page, err := cmd.Flags().GetInt("page")
		if err != nil {
			logrus.Errorf("Unable to get name from cli: %v\n", err)
			return
		}

		finished, err := cmd.Flags().GetBool("finished")
		if err != nil {
			logrus.Errorf("Unable to get name from cli: %v\n", err)
			return
		}

		novel := adapter.CreateNovelAdapter{
			Name:     name,
			Page:     page,
			Finished: finished,
		}

		fmt.Println(novel.Finished)
		os.Exit(0)

		novelController := controller.GetNovelControllerInstance()

		novelController.CliCreate(&novel)
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
	novelCreateCmd.Flags().Bool("finished", false, "flag for defining if the novel/book is still launching")

	novelGetCmd.Flags().String("name", "", "get novel by name")
	novelGetCmd.Flags().String("token", "", "get novel by token")

	novelUpdateCmd.Flags().String("name", "", "get novel by name")
	novelUpdateCmd.Flags().String("token", "", "get novel by name")
	novelUpdateCmd.Flags().Int("page", 0, "page that I'm currently in")
	novelUpdateCmd.Flags().Bool("finished", false, "flag for defining if the novel/book is still launching")

	novelDeleteCmd.Flags().String("token", "", "delete a novel by the token")
}
