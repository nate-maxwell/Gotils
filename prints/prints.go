package prints

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"gotils/terminal"
)

func SprintfLn(formatStr string, args ...string) {
	interfaceArgs := make([]any, len(args))
	for i, v := range args {
		interfaceArgs[i] = v
	}
	msg := fmt.Sprintf(formatStr, interfaceArgs...)
	fmt.Println(msg)
}

// Prints a "-----header-----" to fill the terminal length if a terminal is
// being used for Stdout, otherwise prints 80 columns wide.
func PrintCenteredHeader(header string) {
	width := terminal.GetOutputWidth()
	vis := utf8.RuneCountInString(header)
	spacer := (width - vis) / 2
	side := strings.Repeat("-", spacer)
	SprintfLn("%s%s%s", side, header, side)
}

// Prints "-" repeated to fill the terminal length if a terminal is being used
// for Stdout, otherwise repeats 80 columns wide.
func PrintAsciiLine() {
	width := terminal.GetOutputWidth()
	fmt.Println(strings.Repeat("-", width))
}
