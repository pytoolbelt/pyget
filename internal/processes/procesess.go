package processes

import (
	"fmt"
	"github.com/pytoolbelt/pyget/internal/utils"
	"os"
	"os/exec"
	"time"
)

func ExtractPythonTarball(source, target string) error {
	cmd := exec.Command("tar", "-xvf", source, "-C", target)
	return cmd.Run()
}

// RunPythonConfigureScript runs the ./configure script with common flags and a specified prefix.
func RunPythonConfigureScript(path, prefix string) error {
	cmd := exec.Command("./configure",
		"--enable-optimizations",
		"--with-lto",
		"--enable-shared",
		"--with-system-ffi",
		"--with-computed-gotos",
		fmt.Sprintf("--prefix=%s", prefix),
	)
	cmd.Dir = path
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	env := os.Environ()
	cmd.Env = append(cmd.Env, env...)

	// Create a channel to signal the spinner to stop
	done := make(chan bool)
	go utils.Spinner(100*time.Millisecond, done, "Running configure script ", path)

	err := cmd.Run()

	// Signal the spinner to stop
	done <- true

	if err != nil {
		return fmt.Errorf("error running configure script: %s", err)
	}
	return nil
}

func RunPythonMake(path string) error {
	cmd := exec.Command("make")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	env := os.Environ()
	cmd.Env = append(cmd.Env, env...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running make: %s", err)
	}
	return nil
}

func RunPythonMakeInstall(path string) error {
	cmd := exec.Command("make", "install")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	env := os.Environ()
	cmd.Env = append(cmd.Env, env...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running make install: %s", err)
	}
	return nil
}
