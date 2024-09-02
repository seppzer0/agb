package tool

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// RunCmd executes a specified shell command.
func RunCmd(command string) (string, error) {
	var stdout bytes.Buffer

	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	// process stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	err := cmd.Run()

	return strings.TrimSpace(stdout.String()), err
}

// RunCmdWDir executes a specified shell command in a specified working directory.
func RunCmdWDir(command string, path string) (string, error) {
	var stdout bytes.Buffer

	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	// process stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	cmd.Dir = path
	err := cmd.Run()

	return strings.TrimSpace(stdout.String()), err
}

// RunCmdWDirVerbose executes a specified shell command in a working specified directory with a verbose output.
func RunCmdWDirVerbose(command string, path string) error {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path

	output, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	scanner := bufio.NewScanner(output)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\n%s", scanner.Text())
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		return err
	}

	wg.Wait()

	return cmd.Wait()
}
