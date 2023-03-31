/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/scaf-fold/tools/pkg/gormer"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate models",
	Run:   migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("conf")
	if err != nil {
		panic(err)
	}
	assemble := gormer.NewModelConverter(path)
	assemble.Gen()
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringP("conf", "c", "", "配置文件路径")
}
