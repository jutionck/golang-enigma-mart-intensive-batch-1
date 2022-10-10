package utils

import (
	"bufio"
	"os"
)

func BufioInput() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner
}
