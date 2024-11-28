/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/utils"
	"os"

	"github.com/pytoolbelt/pyget/internal/paths"
	"github.com/spf13/cobra"
)

func initEntrypoint(cmd *cobra.Command, args []string) {
	p, err := paths.NewPyGetPaths("3.13.0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = p.CreateDirs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = utils.AppendToProfile(p.BinDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("PyGet initialized")

}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run:   initEntrypoint,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
