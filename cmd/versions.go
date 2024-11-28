/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/paths"
	"os"

	"github.com/spf13/cobra"
)

func versionsEntrypoint(cmd *cobra.Command, args []string) {
	paths, err := paths.NewPyGetPaths("0.0.0")
	if err != nil {
		fmt.Printf("Error creating paths: %s\n", err)
		os.Exit(1)
	}

	versions, err := paths.FindPythonBinaries()
	if err != nil {
		fmt.Printf("Error finding python binaries: %s\n", err)
		os.Exit(1)
	}

	paths.PrintInstalledPythonVersionsRaw(versions)
}

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List the installed versions of python on this machine",
	Long:  ``,
	Run:   versionsEntrypoint,
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
