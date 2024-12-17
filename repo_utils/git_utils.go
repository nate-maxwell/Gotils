// # Git Repo Utils
//
// * A utility lib for handling operations in the git repo through
// the OS command line.

package repo_utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/chigopher/pathlib"
)

// runGitCommand will run the given git commands in the given directory.
//
// Args:
//
//	dir(pathlib.Path): The path to the repo to run the commands in.
//	args(...string): All commands to run. 'git' is not needed as a command.
//
// Returns:
//
//	error: The command execution error, if any.
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
//
// Args:
//
//	dir(pathlib.Path): The git repo path to navigate to.
//
// Returns:
//
// error: The os.Chdir or exec.Command error that could have arrisen.
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
//
// Args:
//
//	dir(pathlib.Path): The git repo path to navigate to.
//	branch(string): The name of the git branch to switch to.
//
// Returns:
//
//	error: The os.Chdir or exec.Command error that could have arrisen.
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
