package utils

import (
	"fmt"
	"github.com/blang/semver/v4"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func GetMacOSVersion() (string, error) {
	cmd := exec.Command("sw_vers", "-productVersion")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Spinner(delay time.Duration, done chan bool, msg, url string) {
	for {
		select {
		case <-done:
			fmt.Printf("\rDownload complete! %s\n", url)
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c %s %s", r, msg, url)
				time.Sleep(delay)
			}
		}
	}
}

func VersionIsValid(version string) bool {
	_, err := semver.Parse(version)
	if err != nil {
		return false
	}
	return true
}

// AppendToProfile appends a given path to the user's terminal profile.
func AppendToProfile(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	shell := os.Getenv("SHELL")
	var profilePath string

	switch {
	case strings.Contains(shell, "bash"):
		profilePath = filepath.Join(usr.HomeDir, ".bash_profile")
	case strings.Contains(shell, "zsh"):
		profilePath = filepath.Join(usr.HomeDir, ".zshrc")
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	file, err := os.OpenFile(profilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open profile file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\nexport PATH=\"%s:$PATH\"\n", path))
	if err != nil {
		return fmt.Errorf("failed to write to profile file: %w", err)
	}

	return nil
}
