/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/paths"
	"github.com/pytoolbelt/pyget/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

func globalEntrypoint(cmd *cobra.Command, args []string) {
	if !utils.VersionIsValid(globalVersionVar) {
		fmt.Println("Invalid version")
		os.Exit(1)
	}

	paths, err := paths.NewPyGetPaths(globalVersionVar)
	if err != nil {
		fmt.Printf("Error creating paths: %s\n", err)
		os.Exit(1)
	}

	if !paths.VersionInstalled() {
		fmt.Printf("Version %s is not installed\n", globalVersionVar)
		os.Exit(1)
	}

	err = paths.RemoveGlobalSymlinksIfExists()
	if err != nil {
		fmt.Printf("Error removing symlinks: %s\n", err)
		os.Exit(1)
	}

	err = paths.CreateGlobalPythonSymlinks()
	if err != nil {
		fmt.Printf("Error creating symlinks: %s\n", err)
		os.Exit(1)
	}
}

// globalCmd represents the global command
var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Set the global version of python to use",
	Long:  ``,
	Run:   globalEntrypoint,
}

var globalVersionVar string

func init() {
	rootCmd.AddCommand(globalCmd)
	globalCmd.Flags().StringVarP(&globalVersionVar, "version", "v", "0.0.0", "The version of python to set as the global version")
	_ = globalCmd.MarkFlagRequired("version")
}
