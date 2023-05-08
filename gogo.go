package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filesWithDoubles := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			filesWithDoubles[strings.Split(line, "___")[0]] = true
		}
	}
	for ind := range filesWithDoubles {
		fmt.Println(ind)
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[fmt.Sprintf("%s___%s", f.Name(), input.Text())]++
	}
	// Примечание: игнорируем потенциальные
	// ошибки из input.Err()
}
