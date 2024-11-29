package cmd

import (
	"fmt"
	"os"
)

var (
	purple = "\033[1;35m"
	orange = "\033[1;33m"
	white  = "\033[1;37m"
)

func resetTerminal() {
	fmt.Fprint(os.Stdout, "\033[H\033[2J") // Clear screen (opsional)
}
