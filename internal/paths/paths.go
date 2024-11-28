package paths

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strings"
)

type PyGetPaths struct {
	Home                 string
	PyGetRoot            string
	InstallRoot          string
	Version              string
	InstallDir           string
	ExtractionDir        string
	SourceDir            string
	BinDir               string
	PythonExecutable     string
	PythonGlobalSymlink  string
	Python3GlobalSymlink string
	PythonVersionSymlink string
	Pip3Executable       string
	Pip3GlobalSymlink    string
	PipGlobalSymlink     string
}

func NewPyGetPaths(version string) (*PyGetPaths, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &PyGetPaths{
		Home:                 usr.HomeDir,
		PyGetRoot:            usr.HomeDir + "/.pyget",
		InstallRoot:          usr.HomeDir + "/.pyget/versions",
		Version:              version,
		ExtractionDir:        usr.HomeDir + "/.pyget/versions/extraction/" + version,
		SourceDir:            usr.HomeDir + "/.pyget/versions/extraction/" + version + "/Python-" + version,
		InstallDir:           usr.HomeDir + "/.pyget/versions/installed/" + version,
		BinDir:               usr.HomeDir + "/.pyget/bin",
		PythonExecutable:     usr.HomeDir + "/.pyget/versions/installed/" + version + "/bin/python3",
		PythonGlobalSymlink:  usr.HomeDir + "/.pyget/bin/python",
		Python3GlobalSymlink: usr.HomeDir + "/.pyget/bin/python3",
		PythonVersionSymlink: usr.HomeDir + "/.pyget/bin/python" + version,
		Pip3Executable:       usr.HomeDir + "/.pyget/versions/installed/" + version + "/bin/pip3",
		Pip3GlobalSymlink:    usr.HomeDir + "/.pyget/bin/pip3",
		PipGlobalSymlink:     usr.HomeDir + "/.pyget/bin/pip",
	}, nil
}

func (p *PyGetPaths) CreateDirs() error {
	dirs := []string{p.PyGetRoot, p.InstallDir, p.ExtractionDir, p.InstallDir, p.BinDir}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PyGetPaths) FindPythonBinaries() ([]string, error) {
	var binaries []string

	err := filepath.Walk(p.InstallRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "python3" {
			parts := strings.Split(path, "/")
			binaries = append(binaries, parts[len(parts)-3])
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	sort.Strings(binaries)
	return binaries, nil
}

func (p *PyGetPaths) PrintInstalledPythonVersionsTable(binaries []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Installed Python Versions"})

	for _, binary := range binaries {
		table.Append([]string{binary})
	}

	table.Render()
}

func (p *PyGetPaths) PrintInstalledPythonVersionsRaw(binaries []string) {
	fmt.Println("Installed Python Versions")
	for _, binary := range binaries {
		fmt.Println(binary)
	}
}

func (p *PyGetPaths) CreateGlobalPythonSymlinks() error {
	sources := []string{p.PythonGlobalSymlink, p.Python3GlobalSymlink, p.PythonVersionSymlink}

	// create the python executable symlinks
	for _, source := range sources {
		err := os.Symlink(p.PythonExecutable, source)
		if err != nil {
			return err
		}
	}

	sources = []string{p.PipGlobalSymlink, p.Pip3GlobalSymlink}
	for _, source := range sources {
		err := os.Symlink(p.Pip3Executable, source)
		if err != nil {
			return err
		}
	}
	return nil
}
