package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/term"

	"gotils/globals"
)

// Clears the terminal of any current output.
func ClearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Returns the current terminal width if it can be found else 80.
func GetOutputWidth() int {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil {
			return width
		}
	}
	return globals.DEFAULT_TERMINAL_W
}

// Prints a progress bar of the given length, filled to the given percent.
func DrawProgressBar(len int, percent float64) {
	filledLen := int(float64(len) * percent)

	bar := strings.Repeat("█", filledLen) + strings.Repeat(" ", len-filledLen)

	msg := fmt.Sprintf("|%s| %.2f%%", bar, percent*100)
	fmt.Println(msg)
}
