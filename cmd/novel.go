/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/controller"
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

		novelController := controller.GetNovelControllerInstance()

		novelController.CliCreate(&novel)
	},
}

var novelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			logrus.Errorf("Unable to get id from cli: %v\n", err)
			return
		}

		novelController := controller.GetNovelControllerInstance()
		novelController.CliDelete(id)
	},
}

var novelUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			logrus.Errorf("Unable to get id from cli: %v\n", err)
		}

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

		novel := adapter.UpdateNovelAdapter{
			Id:       id,
			Name:     name,
			Page:     page,
			Finished: finished,
		}

		novelController := controller.GetNovelControllerInstance()
		novelController.CliUpdate(&novel)
	},
}

var novelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "create a new book log",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			logrus.Errorf("Unable to get id from cli: %v\n", err)
			return
		}

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			logrus.Errorf("Unable to get all from cli: %v\n", err)
			return
		}

		novelController := controller.GetNovelControllerInstance()

		if all {
			novelController.CliGetAll()
			return
		}
		novelController.CliGet(id)
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

	novelGetCmd.Flags().String("id", "", "get novel by id")
	novelGetCmd.Flags().Bool("all", false, "get every novel")

	novelUpdateCmd.Flags().String("name", "", "get novel by name")
	novelUpdateCmd.Flags().String("id", "", "get novel by name")
	novelUpdateCmd.Flags().Int("page", 0, "page that I'm currently in")
	novelUpdateCmd.Flags().Bool("finished", false, "flag for defining if the novel/book is still launching")

	novelDeleteCmd.Flags().String("id", "", "delete a novel by the token")
}
