package python

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/utils"
	"time"
)

type Downloader interface {
	Download() (string, error)
	GetDownloadURL() string
}

type URLMetadata struct {
	RequiredInstaller string
	DownloadFile      string
}

type PyDownloader struct {
	DownloadTarget string
	Version        string
	meta           *URLMetadata
}

func NewPythonDownloader(target, version string) (*PyDownloader, error) {
	return &PyDownloader{
		DownloadTarget: target,
		Version:        version,
		meta:           NewURLMetadata(version),
	}, nil
}

func NewURLMetadata(version string) *URLMetadata {
	return &URLMetadata{
		RequiredInstaller: "python-source",
		DownloadFile:      fmt.Sprintf("Python-%s.tgz", version),
	}
}

func (pd *PyDownloader) GetDownloadPath() string {
	return pd.DownloadTarget + "/" + pd.meta.DownloadFile
}

func (meta *URLMetadata) getDownloadURL() string {
	return fmt.Sprintf("https://www.python.org/ftp/python/%s/", meta.DownloadFile)
}

func (meta *URLMetadata) GetPythonTarballURL(version string) string {
	return fmt.Sprintf("https://www.python.org/ftp/python/%s/Python-%s.tgz", version, version)
}

func (pd *PyDownloader) Download() error {
	url := pd.meta.GetPythonTarballURL(pd.Version)
	if url == "" {
		return fmt.Errorf("no download URL found")
	}

	done := make(chan bool)
	go utils.Spinner(100*time.Millisecond, done, "Fetching python tarball ", url)

	err := utils.DownloadFile(url, pd.GetDownloadPath())
	if err != nil {
		done <- true
		return err
	}
	done <- true
	return nil
}

//
//func GetVersionToInstall(version string) (*URLMetadata, error) {
//	os := runtime.GOOS
//	arch := runtime.GOARCH
//	var file string
//
//	switch os {
//
//	case "darwin":
//		macos, err := utils.GetMacOSVersion()
//		if err != nil {
//			return nil, err
//		}
//		file = "python-" + version + "-macosx" + macos + ".pkg"
//
//	case "linux":
//		if arch == "amd64" {
//			file = "python-" + version + "-linux-x86_64.tar.xz"
//		} else if arch == "arm64" {
//			file = "python-" + version + "3.9.7-linux-aarch64.tar.xz"
//		}
//
//	case "windows":
//		if arch == "amd64" {
//			file = "python-" + version + "-amd64.exe"
//		} else if arch == "386" {
//			file = "python-" + version + ".exe"
//		}
//	default:
//		return nil, fmt.Errorf("unsupported OS or architecture")
//	}
//	return &URLMetadata{
//		RequiredInstaller: os,
//		DownloadFile:      file,
//	}, nil
//}
