package python

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/utils"
	"time"
)

type PyDownloader struct {
	DownloadTarget string
	Version        string
	FileName       string
}

func NewPythonDownloader(target, version string) (*PyDownloader, error) {
	return &PyDownloader{
		DownloadTarget: target,
		Version:        version,
		FileName:       fmt.Sprintf("Python-%s.tgz", version),
	}, nil
}

func (pd *PyDownloader) GetDownloadPath() string {
	return pd.DownloadTarget + "/" + pd.FileName
}

func (pd *PyDownloader) GetPythonTarballURL() string {
	return fmt.Sprintf("https://www.python.org/ftp/python/%s/Python-%s.tgz", pd.Version, pd.Version)
}

func (pd *PyDownloader) Download() error {
	url := pd.GetPythonTarballURL()
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
