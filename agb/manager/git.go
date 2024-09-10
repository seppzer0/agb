package manager

import (
	cerror "agb/error"
	"agb/tool"
	"fmt"
)

// GitManager manages Git-related operations.
type GitManager struct{}

func NewGitManager() *GitManager {
	return &GitManager{}
}

// Clone downloads specified repository.
func (gm *GitManager) Clone(url string, path string, shallow bool) error {
	var depth_flag string

	if shallow {
		depth_flag = " --depth 1"
	} else {
		depth_flag = ""
	}

	command := fmt.Sprintf("git clone%s %s %s", depth_flag, url, path)

	if out, err := tool.RunCmd(command); err != nil {
		return cerror.ErrCommandRun{Command: command, Output: out}
	}

	return nil
}

// Reset hard resets the state of the cloned git repository.
func (gm *GitManager) Reset(path string) error {
	if _, err := tool.RunCmdWDir("git clean -fdx", path); err != nil {
		return err
	}

	if _, err := tool.RunCmdWDir("git reset --soft", path); err != nil {
		return err
	}

	return nil
}
