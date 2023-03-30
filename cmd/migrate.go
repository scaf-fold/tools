/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	xconvert "github.com/scaf-fold/tools/xconverter"
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
	assemble := xconvert.NewModelConverter(path)
	assemble.Gen()
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringP("conf", "c", "", "配置文件路径")
}
