package tool

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

// RunCmd executes shell command.
func RunCmd(command string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = cmd.Stdout

	var wg sync.WaitGroup
	wg.Add(1)

	scanner := bufio.NewScanner(out)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\n%s", scanner.Text())
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		return "", err
	}

	wg.Wait()

	// convert output into string
	fout, err := io.ReadAll(out)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(fout)), cmd.Wait()
}

// RunCmdWDir executes shell command in specified directory.
func RunCmdWDir(command string, path string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path

	out, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = cmd.Stdout

	var wg sync.WaitGroup
	wg.Add(1)

	scanner := bufio.NewScanner(out)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\n%s", scanner.Text())
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		return "", err
	}

	wg.Wait()

	// convert output into string
	fout, err := io.ReadAll(out)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(fout)), cmd.Wait()
}

// RunCmdQuiet executes shell command without printing it.
func RunCmdQuiet(command string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// RunCmdWDirQuiet executes shell command in specified directory without printing it.
func RunCmdWDirQuiet(command string, path string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
