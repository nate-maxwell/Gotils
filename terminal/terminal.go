package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Clears the terminal of any current output.
func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Prints a progress bar of the given length, filled to the given percent.
func DrawProgressBar(len int, percent float64) {
	filledLen := int(float64(len) * percent)

	bar := strings.Repeat("█", filledLen) + strings.Repeat(" ", len-filledLen)

	msg := fmt.Sprintf("|%s| %.2f%%", bar, percent*100)
	fmt.Println(msg)
}
