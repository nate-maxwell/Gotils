package terminal_utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

// Prints a progress bar of the given length, filled to the given percent.
//
// Args:
//
//	len(int): How many characters long the bar should be.
//	percent(float64): What percentage of the bar is filled.
func DrawProgressBar(len int, percent float64) {
	filledLen := int(float64(len) * percent)

	bar := strings.Repeat("█", filledLen) + strings.Repeat(" ", len-filledLen)

	msg := fmt.Sprintf("|%s| %.2f%%", bar, percent*100)
	fmt.Println(msg)
}
