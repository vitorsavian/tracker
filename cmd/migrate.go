/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/internal/env"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
	},
}

var migrateUpCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate up called")

		env.SetEnv()

		origin, err := os.Getwd()
		if err != nil {
			logrus.Fatalf("Error encountered while getting main directory: %v", err)
		}

		m, err := migrate.New(
			fmt.Sprintf("%s%s", "file://", filepath.Join(origin, "internal", "database", "migrations")),
			os.Getenv("DATABASE_URL"))

		if err != nil {
			logrus.Fatal(err)
		}

		if err := m.Up(); err != nil {
			logrus.Fatal(err)
		}

	},
}

var migrateDownCmd = &cobra.Command{
	Use: "down",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate down called")

		origin, err := os.Getwd()
		if err != nil {
			logrus.Fatalf("Error encountered while getting main directory: %v", err)
		}

		m, err := migrate.New(
			fmt.Sprintf("%s%s", "file://", filepath.Join(origin, "internal", "database", "migrations")),
			"postgres://localhost:5432/database?sslmode=enable")
		if err != nil {
			logrus.Fatal(err)
		}

		if err := m.Down(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
