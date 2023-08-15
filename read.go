package out2json

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return strings.Join(input, "\n")
}
