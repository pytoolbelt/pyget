/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/processes"
	"github.com/pytoolbelt/pyget/internal/utils"
	"os"

	"github.com/pytoolbelt/pyget/internal/downloaders/python"
	"github.com/pytoolbelt/pyget/internal/paths"
	"github.com/spf13/cobra"
)

func installEntrypoint(cmd *cobra.Command, args []string) {

	if !utils.VersionIsValid(installVersion) {
		fmt.Println("Invalid version number")
		os.Exit(1)
	}

	pygetPaths, err := paths.NewPyGetPaths(installVersion)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = pygetPaths.CreateDirs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	downloader, err := python.NewPythonDownloader(pygetPaths.InstallDir, pygetPaths.Version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = downloader.Download()
	if err != nil {
		fmt.Println(err)
		err = pygetPaths.RemoveDirs()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	err = processes.ExtractPythonTarball(downloader.GetDownloadPath(), pygetPaths.ExtractionDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = paths.RemoveTarball(downloader.GetDownloadPath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonConfigureScript(pygetPaths.SourceDir, pygetPaths.InstallDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonMake(pygetPaths.SourceDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processes.RunPythonMakeInstall(pygetPaths.SourceDir)
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

var installVersion string

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringVarP(&installVersion, "version", "v", "", "Python version to install")
}

//https://www.python.org/ftp/python/3.12.7/python-3.12.7-macos11.pkg
// https://www.python.org/ftp/python/3.11.4/python-3.11.4-macos11.pkg
// https://www.python.org/ftp/python/3.9.0/python-3.9.0-macosx10.9.pkg
