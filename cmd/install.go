/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/processes"
	"os"

	"github.com/pytoolbelt/pyget/internal/downloaders/python"
	"github.com/pytoolbelt/pyget/internal/paths"
	"github.com/spf13/cobra"
)

func installEntrypoint(cmd *cobra.Command, args []string) {
	paths, err := paths.NewPyGetPaths("3.11.10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = paths.CreateDirs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	downloader, err := python.NewPythonDownloader(paths.InstallDir, paths.Version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = downloader.Download()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.ExtractPythonTarball(downloader.GetDownloadPath(), paths.ExtractionDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonConfigureScript(paths.SourceDir, paths.InstallDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonMake(paths.SourceDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonMakeInstall(paths.SourceDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// https://www.python.org/ftp/python/3.11.10/Python-3.11.10.tgz
// installCmd represents the installation command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "",
	Long:  ``,
	Run:   installEntrypoint,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

//https://www.python.org/ftp/python/3.12.7/python-3.12.7-macos11.pkg
// https://www.python.org/ftp/python/3.11.4/python-3.11.4-macos11.pkg
// https://www.python.org/ftp/python/3.9.0/python-3.9.0-macosx10.9.pkg
