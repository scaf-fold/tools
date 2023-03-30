/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "octopus-tools",
	Short: "八爪鱼",
	Long:  `链接地址直接启动gin 服务`,
}

func AppRun(app func(cmd *cobra.Command, args []string)) {
	rootCmd.Run = app
	execute()
}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
