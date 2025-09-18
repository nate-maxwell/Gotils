package repo_utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/chigopher/pathlib"
)

// runGitCommand will run the given git commands in the given directory.
func runGitCommand(dir *pathlib.Path, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir.String()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %s\noutput: %s", err, output)
	}
	fmt.Printf("Commander output: %s\n", output)
	return nil
}

// PullGitRepo navigates to teh specified directory and runs 'git pull'.
func PullGitRepo(dir *pathlib.Path) error {
	err := os.Chdir(dir.String())
	if err != nil {
		return fmt.Errorf("PullGitRepo failed to change directory: %v", err)
	}

	err = runGitCommand(dir, "pull")
	if err != nil {
		return fmt.Errorf("PullGitRepo failed to pull: %v", err)
	}

	return nil
}

// ChangeGitBranch navigates to teh specified directory and runs 'git checkout <branch>'.
func ChangeGitBranch(dir *pathlib.Path, branch string) error {
	err := os.Chdir(dir.String())
	if err != nil {
		return fmt.Errorf("ChangeGitBranch failed to change directory: %v", err)
	}

	err = runGitCommand(dir, "checkout", branch)
	if err != nil {
		return fmt.Errorf("ChangeGitBranch failed to checkout branch: %v", err)
	}

	return nil
}
